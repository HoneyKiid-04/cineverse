package database

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	psmigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Init() (*gorm.DB, error) {
	// Configure Viper to read from .env file
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading .env file:", err)
	}

	// Construct database connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PORT"),
	)

	// Open database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

func MigrateUp(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	return executeMigration(sqlDB, "up")
}

func MigrateDown(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	return executeMigration(sqlDB, "down")
}

func executeMigration(db *sql.DB, direction string) error {
	// Get absolute path to migrations directory
	migrationsPath, err := filepath.Abs("migrations")
	if err != nil {
		return fmt.Errorf("failed to get migrations path: %w", err)
	}

	// Convert path to URL-friendly format
	migrationsURL := "file://" + filepath.ToSlash(migrationsPath)

	// Create postgres driver instance
	driver, err := psmigrate.WithInstance(db, &psmigrate.Config{})
	if err != nil {
		return fmt.Errorf("failed to create postgres driver: %w", err)
	}

	// Initialize migrate instance
	m, err := migrate.NewWithDatabaseInstance(
		migrationsURL,
		viper.GetString("DB_NAME"),
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}
	defer m.Close()

	// Execute migration based on direction
	var migrationErr error
	if direction == "up" {
		migrationErr = m.Up()
	} else {
		migrationErr = m.Down()
	}

	// Handle migration errors
	if migrationErr != nil && migrationErr != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", migrationErr)
	}

	log.Printf("Successfully executed %s migrations", direction)
	return nil
}

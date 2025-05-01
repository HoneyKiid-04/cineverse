package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	// Get SQL content from up migration file
	content, err := os.ReadFile(filepath.Join("migrations", "000001_init_schema.up.sql"))
	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}

	// Execute the SQL migration
	if err := db.Exec(string(content)).Error; err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	log.Printf("Successfully executed UP migration: %s", "000001_init_schema.up.sql")
	return nil
}

// MigrateDown rolls back all migrations
func MigrateDown(db *gorm.DB) error {
	// Get SQL content from down migration file
	content, err := os.ReadFile(filepath.Join("migrations", "000001_init_schema.down.sql"))
	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}

	// Execute the SQL rollback
	if err := db.Exec(string(content)).Error; err != nil {
		return fmt.Errorf("failed to execute rollback: %w", err)
	}

	log.Printf("Successfully executed DOWN migration: %s", "000001_init_schema.down.sql")
	return nil
}

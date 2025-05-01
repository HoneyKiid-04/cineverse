package main

import (
	"cineverse/internal/database"
	"cineverse/internal/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// Configure Viper to read from .env file
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading .env file:", err)
	}

	db, err := database.Init()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	if err := database.MigrateUp(db); err != nil {
		log.Fatal("Error migrating database:", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Get port from environment variable, default to 8080 if not set
	port := viper.GetString("PORT")
	if port == "" {
		port = "8080"
	}

	// Register routes
	routes.RegisterRoutes(router, db)

	// Start the server
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

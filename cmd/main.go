package main

import (
	"cineverse/internal/database"
	"cineverse/internal/routes"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	db, err := database.Init()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	if err := database.MigrateUp(db); err != nil {
		log.Fatal("Error migrating database:", err)
	}
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// Initialize Gin router

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

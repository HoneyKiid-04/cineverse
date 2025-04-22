package main

import (
	"cineverse/backend/models"
	"cineverse/backend/routes"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"time"
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	models.InitDB()
}

func main() {
	// Set up routes
	routes.SetupRoutes()

	// Start the server
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server starting on port 8080...")
	log.Fatal(server.ListenAndServe())
}

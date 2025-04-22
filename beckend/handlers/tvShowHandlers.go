package main

import (
	"awesomeProject3/handlers"
	"awesomeProject3/models"
	"awesomeProject3/routes"
	"log"
	"net/http"
)

func main() {
	// Initialize the database
	models.InitDB()

	// Set up routes
	routes.SetupRoutes()

	// Start server
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

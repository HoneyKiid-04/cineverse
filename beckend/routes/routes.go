package routes

import (
	"awesomeProject3/handlers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/api/movies", handlers.GetMoviesHandler)
	http.HandleFunc("/api/movies/", handlers.GetMovieHandler)
	http.HandleFunc("/api/tv-shows", handlers.GetTVShowsHandler)
	// Additional routes...
}

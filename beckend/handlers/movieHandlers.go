package handlers

import (
	"awesomeProject3/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func GetMoviesHandler(w http.ResponseWriter, r *http.Request) {
	db := models.GetDB()
	rows, err := db.Query("SELECT * FROM movies")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var movie models.Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Year, &movie.Genre, &movie.Poster, &movie.Rating); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		movies = append(movies, movie)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	// Handle creating a movie in the database
}

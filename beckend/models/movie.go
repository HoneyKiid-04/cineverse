package models

type Movie struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Year   int     `json:"year"`
	Genre  string  `json:"genre"`
	Poster string  `json:"poster"`
	Rating float64 `json:"rating"`
}

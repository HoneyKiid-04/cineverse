package models

type TVShow struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Poster string  `json:"poster"`
	Rating float64 `json:"rating"`
}

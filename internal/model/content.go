package model

import (
	"time"

	"gorm.io/gorm"
)

type ContentType string

const (
	MovieType ContentType = "movie"
	ShowType  ContentType = "show"
)

type Content struct {
	gorm.Model
	Title       string      `json:"title" gorm:"not null"`
	Description string      `json:"description"`
	Type        ContentType `json:"type" gorm:"not null"`
	ReleaseDate time.Time   `json:"release_date"`
	Duration    int         `json:"duration"`
	Rating      float32     `json:"rating"`
	Genres      string      `json:"genres"`
	Director    string      `json:"director"`
	Cast        string      `json:"cast"`
	PosterURL   string      `json:"poster_url"`
	// Show specific fields
	Seasons  int `json:"seasons,omitempty"`
	Episodes int `json:"episodes,omitempty"`
}

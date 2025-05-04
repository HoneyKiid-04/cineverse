package service

import (
	"cineverse/internal/model"
	"cineverse/internal/repository"
	"errors"
	"time"
)

// Request structs
type CreateContentInput struct {
	Title       string            `json:"title" binding:"required"`
	Description string            `json:"description"`
	Type        model.ContentType `json:"type" binding:"required"`
	ReleaseDate time.Time         `json:"release_date"`
	Duration    int               `json:"duration"`
	Rating      float32           `json:"rating"`
	Genres      string            `json:"genres"`
	Director    string            `json:"director"`
	Cast        string            `json:"cast"`
	PosterURL   string            `json:"poster_url"`
	Seasons     int               `json:"seasons,omitempty"`
	Episodes    int               `json:"episodes,omitempty"`
}

type UpdateContentInput struct {
	Title       *string            `json:"title"`
	Description *string            `json:"description"`
	Type        *model.ContentType `json:"type"`
	ReleaseDate *time.Time         `json:"release_date"`
	Duration    *int               `json:"duration"`
	Rating      *float32           `json:"rating"`
	Genres      *string            `json:"genres"`
	Director    *string            `json:"director"`
	Cast        *string            `json:"cast"`
	PosterURL   *string            `json:"poster_url"`
	Seasons     *int               `json:"seasons,omitempty"`
	Episodes    *int               `json:"episodes,omitempty"`
}

type ListContentInput struct {
	Page     int `json:"page" binding:"required,min=1"`
	PageSize int `json:"page_size" binding:"required,min=1,max=100"`
}

// Response structs
type ContentResponse struct {
	Content *model.Content `json:"content"`
}

type ContentListResponse struct {
	Contents []model.Content `json:"contents"`
}

// ContentService handles all content-related operations
type ContentService struct {
	contentRepo *repository.ContentRepository
}

// NewContentService creates a new instance of ContentService
func NewContentService(contentRepo *repository.ContentRepository) *ContentService {
	return &ContentService{
		contentRepo: contentRepo,
	}
}

// Create creates a new content
func (s *ContentService) Create(input CreateContentInput) (*ContentResponse, error) {
	content := &model.Content{
		Title:       input.Title,
		Description: input.Description,
		Type:        input.Type,
		ReleaseDate: input.ReleaseDate,
		Duration:    input.Duration,
		Rating:      input.Rating,
		Genres:      input.Genres,
		Director:    input.Director,
		Cast:        input.Cast,
		PosterURL:   input.PosterURL,
		Seasons:     input.Seasons,
		Episodes:    input.Episodes,
	}

	if err := s.contentRepo.Create(content); err != nil {
		return nil, err
	}

	return &ContentResponse{Content: content}, nil
}

// GetByID retrieves content by ID
func (s *ContentService) GetByID(id uint) (*ContentResponse, error) {
	content, err := s.contentRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("content not found")
	}

	return &ContentResponse{Content: content}, nil
}

// List retrieves all content with pagination
func (s *ContentService) List(input ListContentInput) (*ContentListResponse, error) {
	contents, err := s.contentRepo.List(input.Page, input.PageSize)
	if err != nil {
		return nil, err
	}

	return &ContentListResponse{Contents: contents}, nil
}

// Update updates existing content
func (s *ContentService) Update(id uint, input UpdateContentInput) (*ContentResponse, error) {
	content, err := s.contentRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("content not found")
	}

	// Update only provided fields
	if input.Title != nil {
		content.Title = *input.Title
	}
	if input.Description != nil {
		content.Description = *input.Description
	}
	if input.Type != nil {
		content.Type = *input.Type
	}
	if input.ReleaseDate != nil {
		content.ReleaseDate = *input.ReleaseDate
	}
	if input.Duration != nil {
		content.Duration = *input.Duration
	}
	if input.Rating != nil {
		content.Rating = *input.Rating
	}
	if input.Genres != nil {
		content.Genres = *input.Genres
	}
	if input.Director != nil {
		content.Director = *input.Director
	}
	if input.Cast != nil {
		content.Cast = *input.Cast
	}
	if input.PosterURL != nil {
		content.PosterURL = *input.PosterURL
	}
	if input.Seasons != nil {
		content.Seasons = *input.Seasons
	}
	if input.Episodes != nil {
		content.Episodes = *input.Episodes
	}

	if err := s.contentRepo.Update(content); err != nil {
		return nil, err
	}

	return &ContentResponse{Content: content}, nil
}

// Delete removes content by ID
func (s *ContentService) Delete(id uint) error {
	return s.contentRepo.Delete(id)
}

// GetByType retrieves content by type
func (s *ContentService) GetByType(contentType model.ContentType) (*ContentListResponse, error) {
	contents, err := s.contentRepo.FindByType(contentType)
	if err != nil {
		return nil, err
	}

	return &ContentListResponse{Contents: contents}, nil
}

// SearchByTitle searches content by title
func (s *ContentService) SearchByTitle(title string) (*ContentListResponse, error) {
	contents, err := s.contentRepo.SearchByTitle(title)
	if err != nil {
		return nil, err
	}

	return &ContentListResponse{Contents: contents}, nil
}

// GetByGenre retrieves content by genre
func (s *ContentService) GetByGenre(genre string) (*ContentListResponse, error) {
	contents, err := s.contentRepo.FindByGenre(genre)
	if err != nil {
		return nil, err
	}

	return &ContentListResponse{Contents: contents}, nil
}

// GetByDirector retrieves content by director
func (s *ContentService) GetByDirector(director string) (*ContentListResponse, error) {
	contents, err := s.contentRepo.FindByDirector(director)
	if err != nil {
		return nil, err
	}

	return &ContentListResponse{Contents: contents}, nil
}

// GetTopRated retrieves top-rated content
func (s *ContentService) GetTopRated(limit int) (*ContentListResponse, error) {
	contents, err := s.contentRepo.GetTopRated(limit)
	if err != nil {
		return nil, err
	}

	return &ContentListResponse{Contents: contents}, nil
}

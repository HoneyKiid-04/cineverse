package repository

import (
	"cineverse/internal/model"

	"gorm.io/gorm"
)

type ContentRepository struct {
	db *gorm.DB
}

func NewContentRepository(db *gorm.DB) *ContentRepository {
	return &ContentRepository{db: db}
}

// Create creates a new content in the database
func (r *ContentRepository) Create(content *model.Content) error {
	return r.db.Create(content).Error
}

// FindByID retrieves content by its ID
func (r *ContentRepository) FindByID(id uint) (*model.Content, error) {
	var content model.Content
	err := r.db.First(&content, id).Error
	if err != nil {
		return nil, err
	}
	return &content, nil
}

// List retrieves all content with pagination
func (r *ContentRepository) List(page, pageSize int) ([]model.Content, error) {
	var contents []model.Content
	offset := (page - 1) * pageSize
	err := r.db.Offset(offset).Limit(pageSize).Find(&contents).Error
	return contents, err
}

// Update updates existing content in the database
func (r *ContentRepository) Update(content *model.Content) error {
	return r.db.Save(content).Error
}

// Delete removes content from the database
func (r *ContentRepository) Delete(id uint) error {
	return r.db.Delete(&model.Content{}, id).Error
}

// FindByType retrieves content by its type (movie or show)
func (r *ContentRepository) FindByType(contentType model.ContentType) ([]model.Content, error) {
	var contents []model.Content
	err := r.db.Where("type = ?", contentType).Find(&contents).Error
	return contents, err
}

// SearchByTitle searches content by title (partial match)
func (r *ContentRepository) SearchByTitle(title string) ([]model.Content, error) {
	var contents []model.Content
	err := r.db.Where("title ILIKE ?", "%"+title+"%").Find(&contents).Error
	return contents, err
}

// FindByGenre finds content by genre
func (r *ContentRepository) FindByGenre(genre string) ([]model.Content, error) {
	var contents []model.Content
	err := r.db.Where("genres ILIKE ?", "%"+genre+"%").Find(&contents).Error
	return contents, err
}

// FindByDirector finds content by director
func (r *ContentRepository) FindByDirector(director string) ([]model.Content, error) {
	var contents []model.Content
	err := r.db.Where("director ILIKE ?", "%"+director+"%").Find(&contents).Error
	return contents, err
}

// GetTopRated retrieves top-rated content with limit
func (r *ContentRepository) GetTopRated(limit int) ([]model.Content, error) {
	var contents []model.Content
	err := r.db.Order("rating DESC").Limit(limit).Find(&contents).Error
	return contents, err
}

package routes

import (
	"cineverse/internal/delivery"
	"cineverse/internal/middleware"
	"cineverse/internal/model"
	"cineverse/internal/repository"
	"cineverse/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterContentRoutes(router *gin.Engine, db *gorm.DB) {
	contentRepo := repository.NewContentRepository(db)
	contentService := service.NewContentService(contentRepo)
	contentHandler := delivery.NewContentHandler(contentService)

	content := router.Group("/api/v1/content")
	{
		// Public routes
		content.GET("", contentHandler.List)
		content.GET("/:id", contentHandler.GetByID)
		content.GET("/type/:type", contentHandler.GetByType)
		content.GET("/search", contentHandler.SearchByTitle)

		// Protected routes (require authentication and admin role)
		protected := content.Use(middleware.AuthMiddleware("HoneySecret"))
		admin := protected.Use(middleware.RequireRole(model.ModeratorRole))
		{
			admin.POST("", contentHandler.Create)
			admin.PUT("/:id", contentHandler.Update)
			admin.DELETE("/:id", contentHandler.Delete)
		}
	}
}

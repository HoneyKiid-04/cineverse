package routes

import (
	"cineverse/internal/delivery"
	"cineverse/internal/middleware"
	"cineverse/internal/repository"
	"cineverse/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserProfileRoutes(router *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	profileService := service.NewUserProfileService(userRepo)
	profileHandler := delivery.NewUserProfileHandler(profileService)

	// All routes require authentication
	profile := router.Group("/api/v1/profile").Use(middleware.AuthMiddleware("HoneySecret"))
	{
		profile.GET("", profileHandler.GetProfile)
		profile.PUT("", profileHandler.UpdateProfile)
		profile.PUT("/password", profileHandler.ChangePassword)
		profile.DELETE("", profileHandler.DeleteProfile)
	}
}

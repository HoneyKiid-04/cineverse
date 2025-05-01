package routes

import (
	"cineverse/internal/delivery"
	"cineverse/internal/repository"
	"cineverse/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(router *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo, "HoneySecret")
	authHandler := delivery.NewAuthHandler(authService)

	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}
}

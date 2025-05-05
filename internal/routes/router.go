package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	RegisterAuthRoutes(router, db)
	RegisterContentRoutes(router, db)
}

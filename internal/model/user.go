package model

import (
	"gorm.io/gorm"
)

type Role string

const (
	ModeratorRole Role = "moderator"
	UserRole      Role = "user"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex;not null"`
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"password" gorm:"not null"`
	Role     Role   `json:"role" gorm:"not null"`
}

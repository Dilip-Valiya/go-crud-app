package models

import (
	"gorm.io/gorm"
)

// User defines the user model with email as the primary key
type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"not null" validate:"required"`
	Email string `json:"email" gorm:"primaryKey;unique;not null" validate:"required,email"`
}

package models

import (
	"gorm.io/gorm"
)

// Todo defines the todo model
type Todo struct {
	gorm.Model
	Title       string `json:"title" gorm:"not null;size:50" validate:"required,max=50"`
	Description string `json:"description" gorm:"size:250" validate:"omitempty,max=250"`
	UserID      uint   `json:"userId" gorm:"not null"`
}

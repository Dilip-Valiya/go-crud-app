package utils

import (
	"github.com/go-playground/validator/v10"
)

// Validator instance
var Validate *validator.Validate

// Initialize the validator
func init() {
	Validate = validator.New()
}

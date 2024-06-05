package handlers

import (
	"encoding/json"
	"go-crud-app/models"
	"go-crud-app/repository"
	"go-crud-app/utils"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a temporary struct
	var loginData struct {
		Email string `json:"email" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Check if the email exists
	var existingUser models.User
	if err := repository.DB.Where("email = ?", loginData.Email).First(&existingUser).Error; err != nil {
		http.Error(w, "user does not exists! please sign up first", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("login success!")
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the user fields
	if err := utils.Validate.Struct(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the email already exists
	var existingUser models.User
	if err := repository.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		http.Error(w, "user already exists, please try login", http.StatusConflict)
		return
	}

	result := repository.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

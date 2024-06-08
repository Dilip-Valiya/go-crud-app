package handlers

import (
	"encoding/json"
	"net/http"

	"go-crud-app/models"
	"go-crud-app/repository"
	"go-crud-app/utils"

	"github.com/gorilla/mux"
)

// GetUsers retrieves all users from the database and sends them as a JSON response.
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	result := repository.DB.Find(&users)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUser retrieves a specific user by ID from the database and sends it as a JSON response.
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	result := repository.DB.First(&user, params["id"])
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Email already exists", http.StatusConflict)
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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User

	if err := repository.DB.First(&user, params["id"]).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Decode the request body into a temporary struct
	var updatedData struct {
		Name string `json:"name" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the updated fields
	if err := utils.Validate.Struct(updatedData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the user fields
	user.Name = updatedData.Name

	// Save the updated user
	if err := repository.DB.Save(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// DeleteUser performs a soft delete on a user in the database.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	if err := repository.DB.First(&user, params["id"]).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	if err := repository.DB.Delete(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("User deleted successfully")
}

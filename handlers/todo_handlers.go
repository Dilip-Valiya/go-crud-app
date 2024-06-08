// handlers/todo_handlers.go

package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-crud-app/middleware"
	"go-crud-app/models"
	"go-crud-app/repository"
	"go-crud-app/utils"

	"github.com/gorilla/mux"
)

// GetUsers retrieves all users from the database and sends them as a JSON response.
func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	result := repository.DB.Find(&todos)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// UpdateTodo updates an existing todo
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	// Parse user email from auth headers
	email, err := middleware.ParseUserEmailFromAuthHeader(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Parse the Todo ID from the URL
	params := mux.Vars(r)
	todoID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid Todo ID", http.StatusBadRequest)
		return
	}

	// Find the user by email
	var user models.User
	if err := repository.DB.Where("email = ?", email).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Find the todo by ID and user ID to ensure ownership
	var todo models.Todo
	if err := repository.DB.Where("id = ? AND user_id = ?", todoID, user.ID).First(&todo).Error; err != nil {
		http.Error(w, "Todo not found or you are not authorized to update this todo", http.StatusNotFound)
		return
	}

	// Decode the request body into a Todo struct
	var updatedTodo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the updated todo fields
	if err := utils.Validate.Struct(updatedTodo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the fields
	todo.Title = updatedTodo.Title
	todo.Description = updatedTodo.Description

	// Save the updated todo
	if err := repository.DB.Save(&todo).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// CreateTodo creates a new todo associated with the authenticated user
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	// Parse user email from auth headers
	email, err := middleware.ParseUserEmailFromAuthHeader(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Find the user by email
	var user models.User
	if err := repository.DB.Where("email = ?", email).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Decode the request body into a Todo struct
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the todo fields
	if err := utils.Validate.Struct(todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Ensure the todo is associated with the authenticated user
	todo.UserID = user.ID

	// Create the todo
	result := repository.DB.Create(&todo)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

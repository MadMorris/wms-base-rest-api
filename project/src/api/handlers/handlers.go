package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MadMorris/wms-base-rest-api/src/domain"
)

// CreateUserHandler handles requests to create a new user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Basic input validation
	if user.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	// Process the user (e.g., call service layer)
	// userService.CreateUser(user)
	// Respond with success or error
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

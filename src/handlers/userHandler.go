// src/handlers/userHandler.go
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pageza/vet-app/src/models"
)

var users []models.User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Placeholder response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Get Users"})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Placeholder response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Create User"})
}

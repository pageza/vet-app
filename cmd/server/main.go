package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pageza/vet-app/internal/database"
	"github.com/pageza/vet-app/internal/handler"
	"github.com/pageza/vet-app/internal/user"
	"gorm.io/gorm"
)

func main() {

	// Initialize your database connection (db).
	var db *gorm.DB

	// Set up the user repository and service.
	userRepo := database.NewUserRepository(db)
	userService := user.NewUserService(userRepo)

	// Create a new router
	r := mux.NewRouter()

	// Intialize handlers
	authHandler := handler.NewAuthHandler()

	//Define Routes
	r.HandleFunc("/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/logout", authHandler.Logout).Methods("POST")
	r.HandleFunc("/profile", authHandler.Profile).Methods("GET")
	r.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		// Extract the user ID from the URL.
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			http.Error(w, "User ID is missing", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		// Use userService to get the user by ID.
		user, err := userService.GetUserByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Marshal the user into JSON.
		jsonResponse, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write the JSON response.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	})

	// Bind to port and pass router in
	log.Fatal(http.ListenAndServe(":8080", r))
}

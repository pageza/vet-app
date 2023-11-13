package main

import (
	"log"
	"net/http"

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

	// Initialize the auth handler
	authHandler := handler.NewAuthHandler()

	// Setup the routes with the auth handler and user service
	mux := handler.SetupRoutes(authHandler, userService)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

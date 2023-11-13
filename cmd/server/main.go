package main

import (
	"log"
	"net/http"

	"github.com/pageza/vet-app/internal/database"
	"github.com/pageza/vet-app/internal/handler"
	"github.com/pageza/vet-app/internal/user"
)

func main() {
	// Initialize your database connection.
	database.Init()
	db := database.DB

	// Set up the user repository and service with the database connection.
	userRepo := database.NewUserRepository(db)
	userService := user.NewUserService(userRepo)

	// Initialize the auth handler without passing userService.
	authHandler := handler.NewAuthHandler()

	// Setup the routes with the auth handler and userService.
	mux := handler.SetupRoutes(authHandler, userService)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pageza/vet-app/internal/handler"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Intialize handlers
	authHandler := handler.NewAuthHandler()

	//Define Routes
	r.HandleFunc("/register", authHandler.Register).Methods("POST")
	r.HandleFunc("login", authHandler.Login).Methods("POST")
	r.HandleFunc("/logout", authHandler.Logout).Methods("POST")
	r.HandleFunc("/profile", authHandler.Profile).Methods("GET")

	// Bind to port and pass router in
	log.Fatal(http.ListenAndServe(":8080", r))
}

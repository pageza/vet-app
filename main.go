package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Define a simple handler for the GET method on the '/' route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World with Gorilla Mux!")
	}).Methods("GET")

	// Get port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	log.Printf("Starting server on port %s\n", port)
	// Start the server using the router
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

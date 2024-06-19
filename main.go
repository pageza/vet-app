package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pageza/vet-app/src/db"
	"github.com/pageza/vet-app/src/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize the database connection
	db.InitDB()
	db.InitRedis()

	// Test PostgreSQL connection
	err = db.DB.Ping()
	if err != nil {
		log.Fatalf("PostgreSQL connection failed: %v", err)
	} else {
		fmt.Println("PostgreSQL connected successfully")
	}

	// Test Redis connection
	result, err := db.RedisClient.Ping(db.RedisCtx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	} else {
		log.Printf("Redis connection successful: %v", result)
	}

	// Set up the router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")

	// Start the server
	port := os.Getenv("PORT")
	log.Println(port)
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

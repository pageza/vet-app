package main

import (
	"fmt"
	"log"
	"net/http"
	"os"


	"github.com/gorilla/mux"
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
	sqlDB, err := db.DB.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL DB: %v", err)
	}
	err = sqlDB.Ping()
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
	r.HandleFunc("/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users/{user_id}/responses", handlers.GetResponsesForUser).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	// Define routes for calls
	r.HandleFunc("/calls", handlers.CreateCall).Methods("POST")
	r.HandleFunc("/calls", handlers.GetCalls).Methods("GET")
	r.HandleFunc("/calls/{id}", handlers.GetCall).Methods("GET")
	r.HandleFunc("/calls/{id}", handlers.UpdateCall).Methods("PUT")
	r.HandleFunc("/calls/{id}", handlers.DeleteCall).Methods("DELETE")

	// Define routes for responses
	r.HandleFunc("/calls/{call_id}/responses", handlers.CreateResponse).Methods("POST")
	r.HandleFunc("/calls/{call_id}/responses", handlers.GetResponses).Methods("GET")
	r.HandleFunc("/responses/{id}", handlers.GetResponse).Methods("GET")
	r.HandleFunc("/responses/{id}", handlers.UpdateResponse).Methods("PUT")
	r.HandleFunc("/responses/{id}", handlers.DeleteResponse).Methods("DELETE")

	// Start the server
	port := os.Getenv("PORT")
	log.Println(port)
	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), r))
}

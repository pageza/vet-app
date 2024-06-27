package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "github.com/pageza/vet-app/config"
    "github.com/pageza/vet-app/db"
)

func main() {
    log.Println("Starting application...")

    // Load environment variables from .env file
    config, err := config.LoadConfig(".")
    if err != nil {
        log.Fatalf("could not load config: %v", err)
    }

    log.Printf("Database host: %s", config.DBHost)
    log.Printf("Database port: %d", config.DBPort)
    log.Printf("Database user: %s", config.DBUser)
    log.Printf("Database password: %s", config.DBPassword)
    log.Printf("Database name: %s", config.DBName)

    // Initialize PostgreSQL
    log.Println("Initializing PostgreSQL...")
    db.InitDB(config)

    // Test PostgreSQL connection
    sqlDB, err := db.DB.DB()
    if err != nil {
        log.Fatalf("Failed to get SQL DB: %v", err)
    }
    err = sqlDB.Ping()
    if err != nil {
        log.Fatalf("PostgreSQL connection failed: %v", err)
    } else {
        log.Println("PostgreSQL connected successfully")
    }

    // Initialize Redis
    log.Println("Initializing Redis...")
    db.InitRedis(config)

    // Test Redis connection
    result, err := db.RedisClient.Ping(db.RedisCtx).Result()
    if err != nil {
        log.Fatalf("Could not connect to Redis: %v", err)
    } else {
        log.Printf("Redis connection successful: %v", result)
    }

    // Set up the router
    log.Println("Setting up the router...")
    r := mux.NewRouter()

    // // Define routes
    // r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
    // r.HandleFunc("/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
    // r.HandleFunc("/users/{user_id}/responses", handlers.GetResponsesForUser).Methods("GET")
    // r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    // r.HandleFunc("/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
    // r.HandleFunc("/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

    // // Define routes for calls
    // r.HandleFunc("/calls", handlers.CreateCall).Methods("POST")
    // r.HandleFunc("/calls", handlers.GetCalls).Methods("GET")
    // r.HandleFunc("/calls/{id}", handlers.GetCall).Methods("GET")
    // r.HandleFunc("/calls/{id}", handlers.UpdateCall).Methods("PUT")
    // r.HandleFunc("/calls/{id}", handlers.DeleteCall).Methods("DELETE")

    // // Define routes for responses
    // r.HandleFunc("/calls/{call_id}/responses", handlers.CreateResponse).Methods("POST")
    // r.HandleFunc("/calls/{call_id}/responses", handlers.GetResponses).Methods("GET")
    // r.HandleFunc("/responses/{id}", handlers.GetResponse).Methods("GET")
    // r.HandleFunc("/responses/{id}", handlers.UpdateResponse).Methods("PUT")
    // r.HandleFunc("/responses/{id}", handlers.DeleteResponse).Methods("DELETE")

    // Start the server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // default port if not specified
    }
    log.Printf("Server running on port %s\n", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), r))
}

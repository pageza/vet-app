package main

import (
    "net/http"
    "net/http/httptest"
    "os"
    "testing"
    "time"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "github.com/pageza/vet-app/src/db"
    "github.com/pageza/vet-app/src/handlers"
)

func TestEnvVariables(t *testing.T) {
    err := godotenv.Load(".env")
    if err != nil {
        t.Fatalf("Error loading .env file: %v", err)
    }

    if os.Getenv("DB_HOST") == "" {
        t.Error("DB_HOST not set")
    }

    if os.Getenv("REDIS_ADDR") == "" {
        t.Error("REDIS_ADDR not set")
    }

    if os.Getenv("PORT") == "" {
        t.Error("PORT not set")
    }
}

func TestRouter(t *testing.T) {
    err := godotenv.Load(".env")
    if err != nil {
        t.Fatalf("Error loading .env file: %v", err)
    }

    db.InitDB()
    db.InitRedis()

    if err := db.DB.Raw("SELECT 1").Error; err != nil {
        t.Fatalf("Database initialization failed: %v", err)
    }

    r := mux.NewRouter()
    r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
    r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
    r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

    req, _ := http.NewRequest("GET", "/users", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    if resp.Code != http.StatusOK {
        t.Fatalf("Expected status OK but got %v", resp.Code)
    }
}

func TestServerStart(t *testing.T) {
    err := godotenv.Load(".env")
    if err != nil {
        t.Fatalf("Error loading .env file: %v", err)
    }

    // Change the port to avoid conflict
    testPort := "8081"
    os.Setenv("PORT", testPort)
	
    go main() // Start the server in a goroutine

    // Wait for the server to start
    time.Sleep(2 * time.Second)

    resp, err := http.Get("http://localhost:" + os.Getenv("PORT") + "/users")
    if err != nil {
        t.Fatalf("Failed to start server: %v", err)
    }

    if resp.StatusCode != http.StatusOK {
        t.Fatalf("Expected status OK but got %v", resp.StatusCode)
    }
}

package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/pageza/vet-app/config"
	"github.com/pageza/vet-app/db"
)

func TestMain(m *testing.M) {
	// Setup code here (if any)
	log.Println("Setting up tests...")

	// Execute tests
	code := m.Run()

	// Cleanup code here (if any)
	log.Println("Tests completed.")

	os.Exit(code)
}

func TestLoadConfig(t *testing.T) {
	_, err := config.LoadConfig(".")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}
}

func TestDatabaseConnections(t *testing.T) {
	config, err := config.LoadConfig(".")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Initialize PostgreSQL
	db.InitDB(config.DB)

	// Test PostgreSQL connection
	sqlDB, err := db.DB.DB()
	if err != nil {
		t.Fatalf("Failed to get SQL DB: %v", err)
	}
	err = sqlDB.Ping()
	if err != nil {
		t.Fatalf("PostgreSQL connection failed: %v", err)
	}

	// Initialize Redis
	db.InitRedis(config)

	// Test Redis connection
	result, err := db.RedisClient.Ping(db.RedisCtx).Result()
	if err != nil {
		t.Fatalf("Could not connect to Redis: %v", err)
	} else if result != "PONG" {
		t.Fatalf("Unexpected Redis ping result: %v", result)
	}
}

func TestRedisInitialization(t *testing.T) {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Initialize Redis
	db.InitRedis(cfg)

	// Test Redis connection
	result, err := db.RedisClient.Ping(db.RedisCtx).Result()
	if err != nil {
		t.Fatalf("Could not connect to Redis: %v", err)
	} else if result != "PONG" {
		t.Fatalf("Unexpected Redis ping result: %v", result)
	}
}

func TestHTTPServer(t *testing.T) {
	// Set up the router
	r := mux.NewRouter()

	// Define a simple test route
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// Create a test server
	ts := httptest.NewServer(r)
	defer ts.Close()

	// Make a request to the test server
	resp, err := http.Get(ts.URL + "/test")
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

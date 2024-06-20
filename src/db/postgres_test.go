package db

import (
    "testing"

    "github.com/joho/godotenv"
)

func TestInitDB(t *testing.T) {
    err := godotenv.Load("../../.env")
    if err != nil {
        t.Fatalf("Error loading .env file: %v", err)
    }

    InitDB()
    var result int
    err = DB.Raw("SELECT 1").Scan(&result).Error
    if err != nil {
        t.Fatalf("PostgreSQL connection failed: %v", err)
    }
    if result != 1 {
        t.Fatalf("Unexpected result from database query: %v", result)
    }
}

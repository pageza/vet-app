package db

import (
    "testing"

    "github.com/joho/godotenv"
)

func TestInitRedis(t *testing.T) {
    err := godotenv.Load("../../.env")
    if err != nil {
        t.Fatalf("Error loading .env file: %v", err)
    }

    InitRedis()
    result, err := RedisClient.Ping(RedisCtx).Result()
    if err != nil {
        t.Fatalf("Could not connect to Redis: %v", err)
    }

    if result != "PONG" {
        t.Fatalf("Unexpected response from Redis: %v", result)
    }
}

package db

import (
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/pageza/vet-app/models"
	"github.com/stretchr/testify/assert"
)

func setupIntegration(t *testing.T) {
	SetupDB(t, &models.User{}, &models.Call{}, &models.Response{})
	setupRedis(t)
}

func TestUserCreationAndSession(t *testing.T) {
	setupIntegration(t)

	// Create a user in PostgreSQL
	user := models.User{Name: "John Doe", Email: "john@example.com"}
	result := DB.Create(&user)
	assert.NoError(t, result.Error)
	assert.NotZero(t, user.ID)

	// Store a session in Redis
	sessionKey := fmt.Sprintf("session:%d", user.ID)
	err := RedisClient.Set(RedisCtx, sessionKey, "session_data", 0).Err()
	assert.NoError(t, err)

	// Retrieve the session from Redis
	sessionData, err := RedisClient.Get(RedisCtx, sessionKey).Result()
	assert.NoError(t, err)
	assert.Equal(t, "session_data", sessionData)

	// Delete the session from Redis
	err = RedisClient.Del(RedisCtx, sessionKey).Err()
	assert.NoError(t, err)

	// Ensure the session is deleted
	sessionData, err = RedisClient.Get(RedisCtx, sessionKey).Result()
	assert.Error(t, err)
	assert.Equal(t, redis.Nil, err)
	assert.Empty(t, sessionData)
}

package db

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/pageza/vet-app/config"

	"github.com/go-redis/redis/v8"

)

func TestInitRedis(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	InitRedis(config)

	result, err := RedisClient.Ping(RedisCtx).Result()
	assert.NoError(t, err)
	assert.Equal(t, "PONG", result)
}

func TestInitRedisConnectionFailure(t *testing.T) {
	invalidConfig := config.Config{
		RedisHost: "invalid_host",
		RedisPort: 6379,
	}

	InitRedis(invalidConfig)

	_, err := RedisClient.Ping(RedisCtx).Result()
	assert.Error(t, err)
}

func TestRedisDataInsertionAndRetrieval(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	err = InitRedis(config)
	assert.NoError(t, err)

	// Set a value in Redis
	err = RedisClient.Set(RedisCtx, "test_key", "test_value", 0).Err()
	assert.NoError(t, err)

	// Get the value from Redis
	val, err := RedisClient.Get(RedisCtx, "test_key").Result()
	assert.NoError(t, err)
	assert.Equal(t, "test_value", val)
}

func TestRedisDataExpiry(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	err = InitRedis(config)
	assert.NoError(t, err)

	// Set a value with an expiry time
	err = RedisClient.Set(RedisCtx, "test_key", "test_value", 1*time.Second).Err()
	assert.NoError(t, err)

	// Ensure the value is set
	val, err := RedisClient.Get(RedisCtx, "test_key").Result()
	assert.NoError(t, err)
	assert.Equal(t, "test_value", val)

	// Wait for the key to expire
	time.Sleep(2 * time.Second)

	// Ensure the value has expired
	val, err = RedisClient.Get(RedisCtx, "test_key").Result()
	assert.Error(t, err)
	assert.Equal(t, redis.Nil, err)
	assert.Empty(t, val)
}
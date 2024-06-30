package db

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pageza/vet-app/config"
	"github.com/stretchr/testify/assert"
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
	setupRedis(t)

	// Set a value in Redis
	err := RedisClient.Set(RedisCtx, "test_key", "test_value", 0).Err()
	assert.NoError(t, err)

	// Get the value from Redis
	val, err := RedisClient.Get(RedisCtx, "test_key").Result()
	assert.NoError(t, err)
	assert.Equal(t, "test_value", val)
}

func TestRedisDataExpiry(t *testing.T) {
	setupRedis(t)

	// Set a value with an expiry time
	err := RedisClient.Set(RedisCtx, "test_key", "test_value", 1*time.Second).Err()
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

func TestRedisIncrement(t *testing.T) {
	setupRedis(t)

	// Increment a key
	err := RedisClient.Incr(RedisCtx, "counter").Err()
	assert.NoError(t, err)

	// Increment the key again
	err = RedisClient.Incr(RedisCtx, "counter").Err()
	assert.NoError(t, err)

	// Get the counter value
	val, err := RedisClient.Get(RedisCtx, "counter").Result()
	assert.NoError(t, err)
	assert.Equal(t, "2", val)
}

func TestRedisListOperations(t *testing.T) {
	setupRedis(t)

	// Push values to a list
	err := RedisClient.RPush(RedisCtx, "list", "value1").Err()
	assert.NoError(t, err)
	err = RedisClient.RPush(RedisCtx, "list", "value2").Err()
	assert.NoError(t, err)

	// Pop a value from the list
	val, err := RedisClient.LPop(RedisCtx, "list").Result()
	assert.NoError(t, err)
	assert.Equal(t, "value1", val)

	// Pop the next value from the list
	val, err = RedisClient.LPop(RedisCtx, "list").Result()
	assert.NoError(t, err)
	assert.Equal(t, "value2", val)
}

func TestRedisHashOperations(t *testing.T) {
	setupRedis(t)

	// Set a hash field
	err := RedisClient.HSet(RedisCtx, "hash", "field1", "value1").Err()
	assert.NoError(t, err)

	// Get the hash field value
	val, err := RedisClient.HGet(RedisCtx, "hash", "field1").Result()
	assert.NoError(t, err)
	assert.Equal(t, "value1", val)
}

func TestRedisSetOperations(t *testing.T) {
	setupRedis(t)

	// Add members to a set
	err := RedisClient.SAdd(RedisCtx, "set", "member1").Err()
	assert.NoError(t, err)
	err = RedisClient.SAdd(RedisCtx, "set", "member2").Err()
	assert.NoError(t, err)

	// Check if a member exists in the set
	isMember, err := RedisClient.SIsMember(RedisCtx, "set", "member1").Result()
	assert.NoError(t, err)
	assert.True(t, isMember)

	// Get all members from the set
	members, err := RedisClient.SMembers(RedisCtx, "set").Result()
	assert.NoError(t, err)
	assert.ElementsMatch(t, []string{"member1", "member2"}, members)
}

func TestRedisLargeData(t *testing.T) {
	setupRedis(t)

	// Insert a large value
	largeValue := make([]byte, 1024*1024) // 1 MB
	err := RedisClient.Set(RedisCtx, "large_key", largeValue, 0).Err()
	assert.NoError(t, err)

	// Retrieve the large value
	val, err := RedisClient.Get(RedisCtx, "large_key").Bytes()
	assert.NoError(t, err)
	assert.Equal(t, largeValue, val)
}

func TestConcurrentRedisAccess(t *testing.T) {
	setupRedis(t)

	const numGoroutines = 10
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				key := fmt.Sprintf("counter_%d", i)
				err := RedisClient.Incr(RedisCtx, key).Err()
				assert.NoError(t, err)
			}
		}(i)
	}

	wg.Wait()

	for i := 0; i < numGoroutines; i++ {
		key := fmt.Sprintf("counter_%d", i)
		val, err := RedisClient.Get(RedisCtx, key).Int()
		assert.NoError(t, err)
		assert.Equal(t, 10, val)
	}
}

func setupRedis(t *testing.T) {
	config, err := config.LoadConfig("../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	err = InitRedis(config)
	assert.NoError(t, err)

	// Clear the Redis database before each test
	err = RedisClient.FlushDB(RedisCtx).Err()
	assert.NoError(t, err)
}

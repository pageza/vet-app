package db

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
	RedisCtx    context.Context
)

func InitRedis() {
	RedisCtx = context.Background()

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := RedisClient.Ping(RedisCtx).Result()
	if err != nil {
		panic(err)
	}
}

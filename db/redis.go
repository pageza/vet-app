package db

import (
	"context"
	"fmt"
	"log"
	"time"
	
	"github.com/go-redis/redis/v8"

	"github.com/pageza/vet-app/config"
)

var RedisClient *redis.Client
var RedisCtx = context.Background()

func InitRedis(config config.Config) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.RedisHost, config.RedisPort),
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})

	ctx, cancel := context.WithTimeout(RedisCtx, 5*time.Second)
	defer cancel()

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Printf("Failed to connect to Redis: %v", err)
		return err
	}

	log.Println("Redis connected successfully")
	return nil
}

// src/db/redis.go
package db

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client
var ctx = context.Background()

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}

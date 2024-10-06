package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
)

var (
	redisUrl = os.Getenv("REDIS_URL")
	ctx      = context.Background()
	rdb      *redis.Client
)

func ConnectToRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func RedisSet(key string, value string) {
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		fmt.Printf("Could not set key: %v\n", err)
	}
}

func RedisGet(key string) string {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Printf("Could not get key: %v\n", err)
	}

	return val
}

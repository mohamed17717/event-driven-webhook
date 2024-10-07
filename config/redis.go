package config

import (
	"context"
	"event-driven-webhook/utils"
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
	utils.LogOnError(err, "Could not set key: %v")
}

func RedisGet(key string) string {
	val, err := rdb.Get(ctx, key).Result()
	utils.LogOnError(err, "Could not get key: %v")

	return val
}

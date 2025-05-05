package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var client *redis.Client

func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
}

func GetResult(hash string) (string, error) {
	val, err := client.Get(ctx, hash).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

func SaveResult(hash, result string) {
	client.Set(ctx, hash, result, time.Hour*24) // 24ч TTL
}

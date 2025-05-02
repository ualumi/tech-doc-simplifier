package db

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client
var Ctx = context.Background()

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"), // redis:6379
	})
}

func CheckToken(token string) bool {
	exists, _ := Redis.Exists(Ctx, token).Result()
	return exists == 1
}

func SaveToken(token string, login string, ttl time.Duration) error {
	return Redis.Set(Ctx, token, login, ttl).Err()
}

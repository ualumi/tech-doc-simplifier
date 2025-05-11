package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func NewRedisClient(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{Addr: addr})
}

func GetUserLoginByToken(rdb *redis.Client, token string) (string, error) {
	return rdb.Get(ctx, token).Result()
}

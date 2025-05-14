package redis

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var client *redis.Client

func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	log.Println("Redis client initialized, connecting to redis:6379...")
}

func GetResult(hash string) (string, error) {
	log.Printf("Trying to get result for key: %s", hash)
	val, err := client.Get(ctx, hash).Result()
	if err == redis.Nil {
		log.Printf("Key not found in Redis: %s", hash)
		return "", nil
	} else if err != nil {
		log.Printf("Error retrieving key %s from Redis: %v", hash, err)
		return "", err
	}
	log.Printf("Key %s found in Redis with value: %s", hash, val)
	return val, nil
}

func SaveResult(hash, result string) {
	log.Printf("Saving result to Redis: key=%s, value=%s", hash, result)
	err := client.Set(ctx, hash, result, time.Hour*24).Err()
	if err != nil {
		log.Printf("Error saving key %s to Redis: %v", hash, err)
	} else {
		log.Printf("Successfully saved key %s with 24h TTL", hash)
	}
}

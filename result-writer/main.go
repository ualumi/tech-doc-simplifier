package main

import (
	"log"
)

func main() {
	cfg := Config{
		KafkaBroker: "broker:9092",
		KafkaTopic:  "model_response",
		RedisAddr:   "redis:6379",
		PostgresDSN: "postgresql://postgres:postgres@user-db:5432/users_info?sslmode=disable",
	}

	log.Println("🔌 Connecting to Redis...")
	redisClient := NewRedisClient(cfg.RedisAddr)
	if redisClient == nil {
		log.Fatal("❌ Failed to connect to Redis")
	}
	log.Println("✅ Connected to Redis")

	log.Println("🔌 Connecting to PostgreSQL...")
	db, err := InitDB(cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("❌ Failed to connect to PostgreSQL: %v", err)
	}
	log.Println("✅ Connected to PostgreSQL")

	log.Println("🚀 Starting Kafka consumers...")
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Fatalf("❌ Panic in StartKafkaConsumer: %v", r)
			}
		}()
		StartKafkaConsumer(cfg, redisClient, db)
		log.Println("🛑 Kafka consumer for model_response stopped")
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Fatalf("❌ Panic in StartResultRequestConsumer: %v", r)
			}
		}()
		StartResultRequestConsumer(cfg, redisClient, db)
		log.Println("🛑 Kafka consumer for result_request stopped")
	}()

	log.Println("🟢 result-writer is running and waiting for Kafka messages...")
	log.Println("! result-writer is running and waiting for Kafka messages...")
	select {} // блокировка main
}

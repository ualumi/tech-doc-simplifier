package kafka

import (
	"auth-service/db"
	"context"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

func StartConsumer() {
	broker := os.Getenv("KAFKA_BROKER")
	topic := os.Getenv("KAFKA_TOPIC")

	// Добавляем проверку и логирование
	if topic == "" {
		log.Fatal("KAFKA_TOPIC environment variable is not set")
	}
	if broker == "" {
		log.Fatal("KAFKA_BROKER environment variable is not set")
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		//GroupID: "auth-group",
	})

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("Kafka read error:", err)
			continue
		}

		token := string(msg.Value)
		if db.CheckToken(token) {
			log.Println("Auth token found in Redis:", token)
		} else {
			user, err := db.GetUserByToken(token)
			if err != nil {
				log.Println("Unauthorized:", token)
			} else {
				log.Println("Authorized:", user.Login)
			}
		}
	}
}

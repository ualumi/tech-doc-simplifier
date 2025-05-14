package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	requestWriter  *kafka.Writer
	responseWriter *kafka.Writer
)

func InitProducer() {
	requestWriter = &kafka.Writer{
		Addr:     kafka.TCP("broker:9092"),
		Topic:    "model_requests", // Задаём здесь
		Balancer: &kafka.LeastBytes{},
	}

	responseWriter = &kafka.Writer{
		Addr:     kafka.TCP("broker:9092"),
		Topic:    "model_response",
		Balancer: &kafka.LeastBytes{},
	}
}

func PublishText(key string, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := requestWriter.WriteMessages(ctx, kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
	})
	if err != nil {
		log.Println("Kafka publish error (model_requests):", err)
	}
	return err
}

func PublishSimplifiedResponse(hash string, result string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := responseWriter.WriteMessages(ctx, kafka.Message{
		Key:   []byte(hash),
		Value: []byte(result),
	})
	if err != nil {
		log.Println("Kafka simplified_response publish error:", err)
	}
	return err
}

package kafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

var responseReader *kafka.Reader

func InitConsumer(broker, topic string) {
	responseReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		GroupID:  "api-gateway-response-group",
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
}

func ReadResponse(correlationID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for {
		m, err := responseReader.ReadMessage(ctx)
		if err != nil {
			return "", err
		}
		if string(m.Key) == correlationID {
			return string(m.Value), nil
		}
	}
}

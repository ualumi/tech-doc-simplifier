/*package kafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

var responseReader *kafka.Reader

func InitConsumer(broker, topic string) {
	responseReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{broker},
		Topic:       topic,
		GroupID:     "api-gateway-response-group",
		MinBytes:    10e3,
		MaxBytes:    10e6,
		StartOffset: kafka.FirstOffset,
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

func CloseConsumer() error {
	if responseReader != nil {
		return responseReader.Close()
	}
	return nil
}*/

package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var responseReader *kafka.Reader

func InitConsumer(broker, topic string) {
	log.Printf("[Kafka InitConsumer] Initializing consumer for topic '%s' at broker '%s'\n", topic, broker)

	responseReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		//GroupID:     "api-gateway-response-group",
		MinBytes:    10e3,
		MaxBytes:    10e6,
		StartOffset: kafka.LastOffset,
	})

	log.Println("[Kafka InitConsumer] Consumer successfully initialized")
}

func ReadResponse(correlationID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Printf("[Kafka ReadResponse] Waiting for response with correlationID: %s\n", correlationID)

	for {
		m, err := responseReader.ReadMessage(ctx)
		if err != nil {
			log.Printf("[Kafka ReadResponse] Error while reading message: %v\n", err)
			return "", err
		}

		log.Printf("[Kafka ReadResponse] Received message from topic: key=%s, value=%s\n", string(m.Key), string(m.Value))

		if string(m.Key) == correlationID {
			log.Printf("[Kafka ReadResponse] Found matching response for correlationID: %s\n", correlationID)
			return string(m.Value), nil
		}
	}
}

func CloseConsumer() error {
	if responseReader != nil {
		log.Println("[Kafka CloseConsumer] Closing Kafka consumer")
		return responseReader.Close()
	}
	return nil
}

package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var responseReader *kafka.Reader

// new one
var resultResponseReader *kafka.Reader

func InitConsumer(broker, topic string) {
	log.Printf("[Kafka InitConsumer] Initializing consumer for topic '%s' at broker '%s'\n", topic, broker)

	// Сброс offset до самого последнего вручную
	conn, err := kafka.DialLeader(context.Background(), "tcp", broker, topic, 0)
	if err != nil {
		log.Fatalf("[Kafka InitConsumer] Failed to connect to Kafka leader: %v\n", err)
	}
	lastOffset, err := conn.ReadLastOffset()
	if err != nil {
		log.Fatalf("[Kafka InitConsumer] Failed to read last offset: %v\n", err)
	}
	conn.Close()

	responseReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{broker},
		Topic:       topic,
		GroupID:     "",         // без GroupID — каждый раз читаем заново
		StartOffset: lastOffset, // начинаем с последнего оффсета
		MinBytes:    10e3,
		MaxBytes:    10e6,
	})

	log.Println("[Kafka InitConsumer] Consumer successfully initialized")
}

func ReadResponse(correlationID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
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

		log.Printf("[Kafka ReadResponse] No match for correlationID: %s, got: %s\n", correlationID, string(m.Key))
	}
}

func CloseConsumer() error {
	if responseReader != nil {
		log.Println("[Kafka CloseConsumer] Closing Kafka consumer")
		return responseReader.Close()
	}
	return nil
}

func InitResultResponseConsumer(broker string) {
	const topic = "result_response"
	log.Printf("[Kafka InitResultResponseConsumer] Initializing consumer for topic '%s'\n", topic)

	resultResponseReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		GroupID:  "",
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	log.Println("[Kafka InitResultResponseConsumer] Consumer for result_response initialized")
}

func ReadResultResponse(correlationID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	log.Printf("[Kafka ReadResultResponse] Waiting for response with correlationID: %s\n", correlationID)

	for {
		m, err := resultResponseReader.ReadMessage(ctx)
		if err != nil {
			log.Printf("[Kafka ReadResultResponse] Error while reading message: %v\n", err)
			return "", err
		}

		log.Printf("[Kafka ReadResultResponse] Received message: key=%s, value=%s\n", string(m.Key), string(m.Value))

		if string(m.Key) == correlationID {
			log.Printf("[Kafka ReadResultResponse] Found matching response for correlationID: %s\n", correlationID)
			return string(m.Value), nil
		}
	}
}

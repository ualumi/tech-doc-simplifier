package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func ReadModelResponseOnce(broker, topic, correlationID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.Printf("[Kafka ReadModelResponseOnce] Waiting for model response with correlationID: %s\n", correlationID)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		GroupID:  "read-model-response-" + correlationID,
		MinBytes: 1,
		MaxBytes: 10e6,
	})

	conn, err := kafka.DialLeader(ctx, "tcp", broker, topic, 0)
	if err != nil {
		log.Printf("[Kafka ReadModelResponseOnce] Failed to dial leader: %v\n", err)
		return "", err
	}
	lastOffset, err := conn.ReadLastOffset()
	conn.Close()
	if err != nil {
		log.Printf("[Kafka ReadModelResponseOnce] Failed to get last offset: %v\n", err)
		return "", err
	}
	reader.SetOffset(lastOffset)

	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("[Kafka ReadModelResponseOnce] Read error: %v\n", err)
			return "", err
		}

		log.Printf("[Kafka ReadModelResponseOnce] Received message: key=%s, value=%s\n", string(msg.Key), string(msg.Value))

		if string(msg.Key) == correlationID {
			log.Printf("[Kafka ReadModelResponseOnce] Matched correlationID: %s\n", correlationID)
			return string(msg.Value), nil
		}
	}
}

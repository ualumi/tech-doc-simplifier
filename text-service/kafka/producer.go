/*package kafka

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
		Addr:  kafka.TCP("broker:9092"),
		Topic: "model_requests",
	}

	responseWriter = &kafka.Writer{
		Addr:  kafka.TCP("broker:9092"),
		Topic: "simplified_response",
	}
}

func PublishText(hash string, text string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := requestWriter.WriteMessages(ctx, kafka.Message{
		Key:   []byte(hash),
		Value: []byte(text),
	})
	if err != nil {
		log.Println("Kafka publish error:", err)
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
}*/

package kafka

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

var (
	requestWriter  *kafka.Writer
	responseWriter *kafka.Writer
)

func InitProducer() {
	requestWriter = &kafka.Writer{
		Addr:  kafka.TCP("broker:9092"),
		Topic: "model_requests",
	}

	responseWriter = &kafka.Writer{
		Addr:  kafka.TCP("broker:9092"),
		Topic: "simplified_response",
	}
}

func PublishText(hash string, text string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	payload := map[string]string{
		"text": text,
	}
	valueBytes, err := json.Marshal(payload)
	if err != nil {
		log.Println("JSON marshal error:", err)
		return err
	}

	err = requestWriter.WriteMessages(ctx, kafka.Message{
		Key:   []byte(hash),
		Value: valueBytes,
	})
	if err != nil {
		log.Println("Kafka publish error:", err)
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

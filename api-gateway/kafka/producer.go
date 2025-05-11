package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func InitProducer(broker string, topic string) {
	writer = &kafka.Writer{
		Addr:  kafka.TCP(broker),
		Topic: topic,
	}
}

// CloseProducer closes the Kafka writer to release resources
func CloseProducer() error {
	if writer != nil {
		return writer.Close()
	}
	return nil
}

// Struct for Kafka message payload
type SimplifyRequest struct {
	Text  string `json:"text"`
	Token string `json:"token"`
}

// PublishSimplifyRequest publishes a message with correlationID, text, and token to Kafka
func PublishSimplifyRequest(correlationID, text, token string) error {
	// Create the message object with text and token
	message := SimplifyRequest{
		Text:  text,
		Token: token,
	}

	// Serialize the message to JSON
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// Publish the message to Kafka
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	msg := kafka.Message{
		Key:   []byte(correlationID), // Use the correlation ID as the message key
		Value: messageBytes,          // Send the serialized JSON message
	}

	return writer.WriteMessages(ctx, msg)
}

// ДЛЯ /history
var writerResultRequest *kafka.Writer

// InitResultRequestProducer инициализирует Kafka writer для топика result_request
func InitResultRequestProducer(broker string) {
	writerResultRequest = &kafka.Writer{
		Addr:  kafka.TCP(broker),
		Topic: "result_request",
	}
}

// PublishResultRequest отправляет токен по ключу (correlationID) в топик result_request
func PublishResultRequest(correlationID, token string) error {
	request := map[string]string{
		"token": token,
	}

	messageBytes, err := json.Marshal(request)
	if err != nil {
		return err
	}

	log.Printf("[Kafka PublishResultRequest] Sending request with correlationID: %s, token: %s\n", correlationID, token)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	msg := kafka.Message{
		Key:   []byte(correlationID),
		Value: messageBytes,
	}

	return writerResultRequest.WriteMessages(ctx, msg)
}

func CloseResultRequestProducer() error {
	if writerResultRequest != nil {
		return writerResultRequest.Close()
	}
	return nil
}

/*package kafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func InitProducer(broker string, topic string) {
	writer = &kafka.Writer{
		Addr:  kafka.TCP(broker),
		Topic: topic,
	}
}

// CloseProducer closes the Kafka writer to release resources
func CloseProducer() error {
	if writer != nil {
		return writer.Close()
	}
	return nil
}

func PublishSimplifyRequest(text string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	msg := kafka.Message{
		Value: []byte(text),
	}
	return writer.WriteMessages(ctx, msg)
}*/

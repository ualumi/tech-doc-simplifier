package kafka

import (
	"auth-service/db"
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

type SimplifyRequest struct {
	Text  string `json:"text"`
	Token string `json:"token"`
}

func StartConsumer() {
	broker := os.Getenv("KAFKA_BROKER")
	requestTopic := os.Getenv("KAFKA_TOPIC") // например: model_requests
	responseTopic := "user_responses"
	forwardTopic := "text_request"

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   requestTopic,
		//GroupID:     "auth-service-group-v2",
		StartOffset: kafka.LastOffset,
		MinBytes:    10e3,
		MaxBytes:    10e6,
		MaxWait:     500 * time.Millisecond,
	})

	responseWriter := kafka.Writer{
		Addr:  kafka.TCP(broker),
		Topic: responseTopic,
	}

	forwardWriter := kafka.Writer{
		Addr:  kafka.TCP(broker),
		Topic: forwardTopic,
	}

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Kafka read error:", err)
			continue
		}

		log.Printf("Received message: Key=%s, Value=%s", string(msg.Key), string(msg.Value))

		var request SimplifyRequest
		err = json.Unmarshal(msg.Value, &request)
		if err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}

		log.Printf("Token from message: %s", request.Token)

		var response string
		if db.CheckToken(request.Token) {
			response = "Token valid"

			// Отправляем запрос в text_request с тем же ключом
			err := forwardWriter.WriteMessages(context.Background(), kafka.Message{
				Key:   msg.Key,
				Value: msg.Value,
			})
			if err != nil {
				log.Println("Error forwarding message to text_request:", err)
			} else {
				log.Println("Forwarded message to text_request")
			}
		} else {
			response = "Unauthorized"
		}

		log.Printf("Sending response: %s", response)

		// ОБЯЗАТЕЛЬНО отправляем ответ с тем же ключом, что был у запроса
		err = responseWriter.WriteMessages(context.Background(), kafka.Message{
			Key:   msg.Key,
			Value: []byte(response),
		})
		if err != nil {
			log.Println("Error sending response:", err)
		}
	}
}

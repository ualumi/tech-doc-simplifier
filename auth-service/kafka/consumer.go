package kafka

import (
	"auth-service/db"
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

type SimplifyRequest struct {
	Text  string `json:"text"`
	Token string `json:"token"`
}

func StartConsumer() {
	broker := os.Getenv("KAFKA_BROKER")
	requestTopic := os.Getenv("KAFKA_TOPIC") // например: model_requests
	responseTopic := "user_responses"        // для статуса валидации
	forwardTopic := "text_request"           // новая цель, если токен валиден

	// Чтение запросов
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    requestTopic,
		GroupID:  "auth-service-group",
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	// Ответ о статусе токена
	responseWriter := kafka.Writer{
		Addr:  kafka.TCP(broker),
		Topic: responseTopic,
	}

	// Писатель для повторной отправки текста
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

			// Асинхронно форвардим сообщение в text_request
			go func(originalKey []byte, originalValue []byte) {
				err := forwardWriter.WriteMessages(context.Background(), kafka.Message{
					Key:   originalKey,
					Value: originalValue,
				})
				if err != nil {
					log.Println("Error forwarding message to text_request:", err)
				} else {
					log.Println("Forwarded message to text_request")
				}
			}(msg.Key, msg.Value)

		} else {
			response = "Unauthorized"
		}

		log.Printf("Sending response: %s", response)

		err = responseWriter.WriteMessages(context.Background(), kafka.Message{
			Key:   msg.Key,
			Value: []byte(response),
		})
		if err != nil {
			log.Println("Error sending response:", err)
		}
	}
}

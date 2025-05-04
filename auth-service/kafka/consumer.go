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
	requestTopic := os.Getenv("KAFKA_TOPIC")
	responseTopic := "user_responses" // фиксировано или из env

	// Создание нового читателя с правильной конфигурацией
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    requestTopic,
		GroupID:  "auth-service-group", // ID группы
		MinBytes: 10e3,                 // минимальный размер пакета (в байтах)
		MaxBytes: 10e6,                 // максимальный размер пакета (в байтах)
	})

	// Создание писателя, который будет отправлять ответы
	writer := kafka.Writer{
		Addr:  kafka.TCP(broker),
		Topic: responseTopic,
	}

	for {
		// Чтение сообщения
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Kafka read error:", err)
			continue
		}

		// Логирование сообщения
		log.Printf("Received message: Key=%s, Value=%s", string(msg.Key), string(msg.Value))

		token := string(msg.Value)
		var response string

		// Проверка токена в базе данных
		if db.CheckToken(token) {
			response = "Token valid"
		} else {
			response = "Unauthorized"
		}

		// Логирование ответа
		log.Printf("Sending response: %s", response)

		// Отправка ответа в Kafka топик
		err = writer.WriteMessages(context.Background(), kafka.Message{
			Key:   msg.Key,
			Value: []byte(response),
		})
		if err != nil {
			log.Println("Error sending response:", err)
		}
	}
}

//старій код с ним все работает
/*package kafka

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
}*/

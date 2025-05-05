package kafka

import (
	"context"
	"log"
	"text-service/redis" // Импорт клиента Redis

	"github.com/segmentio/kafka-go"
)

// Подписка на text_request: получает текст, который нужно обработать
func StartUserRequestConsumer(msgChan chan string) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"broker:9092"},
		Topic:   "text_request",
		GroupID: "text-service-user-group",
	})
	go func() {
		defer r.Close()
		for {
			m, err := r.ReadMessage(context.Background())
			if err != nil {
				log.Println("Kafka read error (text_request):", err)
				continue
			}
			log.Printf("Received user request: %s", string(m.Value))
			msgChan <- string(m.Value) // отправляем текст запроса в канал
		}
	}()
}

// Подписка на model_responses: получает готовые ответы и сохраняет в Redis
func StartResponseConsumer() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"broker:9092"},
		Topic:   "model_responses",
		GroupID: "text-service-model-group",
	})
	go func() {
		defer r.Close()
		for {
			m, err := r.ReadMessage(context.Background())
			if err != nil {
				log.Println("Kafka read error (model_responses):", err)
				continue
			}

			key := string(m.Key)
			value := string(m.Value)
			log.Printf("Received model response: key=%s, value=%s", key, value)

			// Сохраняем результат в Redis
			redis.SaveResult(key, value)
		}
	}()
}

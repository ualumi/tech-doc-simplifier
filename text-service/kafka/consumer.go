package kafka

import (
	"context"
	"log"
	"text-service/redis"

	"github.com/segmentio/kafka-go"
)

// StartUserRequestConsumer подписывается на text_request и отправляет полные сообщения в канал
func StartUserRequestConsumer(msgChan chan kafka.Message) {
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
			log.Printf("Received user request: key=%s, value=%s", string(m.Key), string(m.Value))
			msgChan <- m // передаём kafka.Message с Key и Value
		}
	}()
}

// StartResponseConsumer подписывается на model_response и сохраняет результат в Redis
func StartResponseConsumer() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"broker:9092"},
		Topic:   "model_response",
		GroupID: "text-service-model-group",
	})
	go func() {
		defer r.Close()
		for {
			m, err := r.ReadMessage(context.Background())
			if err != nil {
				log.Println("Kafka read error (model_response):", err)
				continue
			}

			key := string(m.Key)
			value := string(m.Value)
			log.Printf("Received model response: key=%s, value=%s", key, value)

			redis.SaveResult(key, value)
		}
	}()
}

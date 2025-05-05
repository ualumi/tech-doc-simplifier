package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

var modelResponseReader *kafka.Reader

func InitModelResponseConsumer(broker, topic string) {
	modelResponseReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		GroupID:  "model-response-group",
		MinBytes: 1,
		MaxBytes: 10e6,
	})
}

func StartModelResponseConsumer(handler func(string)) {
	go func() {
		log.Println("Модельный консьюмер запущен для топика model_response...")
		for {
			msg, err := modelResponseReader.ReadMessage(context.Background())
			if err != nil {
				log.Println("Ошибка чтения model_response:", err)
				continue
			}
			handler(string(msg.Value))
		}
	}()
}

func CloseModelResponseConsumer() error {
	if modelResponseReader != nil {
		return modelResponseReader.Close()
	}
	return nil
}

package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

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
			msgChan <- m
		}
	}()
}

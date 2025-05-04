package kafka

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

//func PublishSimplifyRequest(text string) error {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//
//	msg := kafka.Message{
//		Value: []byte(text),
//	}
//	return writer.WriteMessages(ctx, msg)
//}

func PublishSimplifyRequest(correlationID, text string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	msg := kafka.Message{
		Key:   []byte(correlationID),
		Value: []byte(text),
	}
	return writer.WriteMessages(ctx, msg)
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

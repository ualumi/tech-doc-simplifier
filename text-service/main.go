package main

import (
	"log"
	"text-service/handlers"
	"text-service/kafka"
	"text-service/redis"

	kafkaGo "github.com/segmentio/kafka-go"
)

func main() {
	redis.InitRedis()
	kafka.InitProducer()

	msgChan := make(chan kafkaGo.Message)
	kafka.StartUserRequestConsumer(msgChan)

	go func() {
		for msg := range msgChan {
			go handlers.HandleUserRequest(msg)
		}
	}()

	log.Println("Text-service is running...")
	select {}
}

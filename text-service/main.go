package main

import (
	"log"
	"text-service/handlers"
	"text-service/kafka"
	"text-service/redis"
)

func main() {
	redis.InitRedis()
	kafka.InitProducer()

	go kafka.StartResponseConsumer()

	msgChan := make(chan string)
	kafka.StartUserRequestConsumer(msgChan)

	go func() {
		for msg := range msgChan {
			go handlers.HandleUserRequest(msg)
		}
	}()

	log.Println("Text-service is running...")
	select {}
}

package handlers

import (
	"log"
	"text-service/kafka"
	"text-service/redis"
	"text-service/utils"

	kafkago "github.com/segmentio/kafka-go"
)

func HandleUserRequest(msg kafkago.Message) {
	var hash string
	text := string(msg.Value)

	if len(msg.Key) > 0 {
		hash = string(msg.Key)
	} else {
		hash = utils.HashText(text)
	}

	cached, err := redis.GetResult(hash)
	if err != nil {
		log.Printf("Redis error: %v", err)
		return
	}

	if cached != "" {
		log.Printf("Cache hit. Sending simplified response for hash: %s", hash)
		err := kafka.PublishSimplifiedResponse(hash, cached)
		if err != nil {
			log.Printf("Failed to publish simplified response: %v", err)
		}
	} else {
		log.Printf("Cache miss. Sending to model_requests: %s", text)
		err := kafka.PublishText(hash, text)
		if err != nil {
			log.Printf("Failed to publish model request: %v", err)
		}
	}
}

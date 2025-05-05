package handlers

import (
	"log"
	"text-service/kafka"
	"text-service/redis"
	"text-service/utils"
)

func HandleUserRequest(text string) {
	hash := utils.HashText(text)

	// Проверка Redis
	cached, err := redis.GetResult(hash)
	if err != nil {
		log.Printf("Redis error: %v", err)
		return
	}

	if cached != "" {
		// Если найдено — отправить результат в simplified_response
		log.Printf("Cache hit. Sending simplified response for hash: %s", hash)
		kafka.PublishSimplifiedResponse(hash, cached)
	} else {
		// Если нет — отправить текст в model_requests
		log.Printf("Cache miss. Sending to model_requests: %s", text)
		kafka.PublishText(hash, text)
	}
}

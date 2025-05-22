package handlers

import (
	"encoding/json"
	"log"
	"text-service/kafka"
	"text-service/redis"
	"text-service/utils"

	kafkago "github.com/segmentio/kafka-go"
)

type Request struct {
	Text  string `json:"text"`
	Token string `json:"token"`
}

func HandleUserRequest(msg kafkago.Message) {
	var req Request

	//JSON
	if err := json.Unmarshal(msg.Value, &req); err != nil {
		log.Printf("Failed to parse message: %v", err)
		return
	}

	//токен и текст
	if req.Text == "" || req.Token == "" {
		log.Println("Missing text or token in request. Skipping.")
		return
	}

	//Redis
	redisKey := utils.HashText(req.Text)
	cached, err := redis.GetResult(redisKey)
	if err != nil {
		log.Printf("Redis error: %v", err)
		return
	}

	if cached != "" {
		log.Printf("Cache hit. Sending simplified response for token: %s", req.Token)

		//структура ответа
		response := kafka.KafkaMessage{
			Original: kafka.OriginalPayload{
				Text:  req.Text,
				Token: req.Token,
			},
			Simplified: kafka.SimplifiedPayload{
				Text:  cached,
				Token: req.Token,
			},
		}

		//сериализация
		respBytes, err := json.Marshal(response)
		if err != nil {
			log.Printf("Failed to marshal response: %v", err)
			return
		}

		//ответ
		err = kafka.PublishSimplifiedResponse(string(msg.Key), string(respBytes))
		if err != nil {
			log.Printf("Failed to publish simplified response: %v", err)
		}
	} else {
		log.Printf("Cache miss. Sending to model_requests: %s", req.Text)
		err := kafka.PublishText(string(msg.Key), string(msg.Value))
		if err != nil {
			log.Printf("Failed to publish model request: %v", err)
		}
	}
}

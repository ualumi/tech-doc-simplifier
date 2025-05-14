/*package handlers

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
}*/

/*package handlers

import (
	"encoding/json"
	"log"
	"text-service/kafka"
	"text-service/redis"

	kafkago "github.com/segmentio/kafka-go"
)

type Request struct {
	Text  string `json:"text"`
	Token string `json:"token"`
}

func HandleUserRequest(msg kafkago.Message) {
	var req Request
	err := json.Unmarshal(msg.Value, &req)
	if err != nil {
		log.Printf("Failed to parse incoming message: %v", err)
		return
	}

	text := req.Text
	token := req.Token
	if text == "" {
		log.Println("text is missing in the request")
		return
	}

	cached, err := redis.GetResult(text)
	if err != nil {
		log.Printf("Redis error for text %s: %v", text, err)
		return
	}

	if cached != "" {
		log.Printf("Cache hit. Sending simplified response for token: %s", token)
		err := kafka.PublishSimplifiedResponse(token, cached)
		if err != nil {
			log.Printf("Failed to publish simplified response: %v", err)
		}
	} else {
		log.Printf("Cache miss. Sending to model_requests for token %s: %s", token, req.Text)
		err := kafka.PublishText(token, req.Text)
		if err != nil {
			log.Printf("Failed to publish model request: %v", err)
		}
	}
}*/

/*package handlers

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

	if len(msg.Value) > 0 {
		hash = string(msg.Value)
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
		err := kafka.PublishSimplifiedResponse(string(msg.Key), cached)
		if err != nil {
			log.Printf("Failed to publish simplified response: %v", err)
		}
	} else {
		log.Printf("Cache miss. Sending to model_requests: %s", text)
		err := kafka.PublishText(string(msg.Key), text)
		if err != nil {
			log.Printf("Failed to publish model request: %v", err)
		}
	}
}*/

package handlers

import (
	"encoding/json"
	"log"
	"text-service/kafka"
	"text-service/redis"

	kafkago "github.com/segmentio/kafka-go"
)

type Request struct {
	Text  string `json:"text"`
	Token string `json:"token"`
}

func HandleUserRequest(msg kafkago.Message) {
	var req Request

	// Распаковываем JSON из Kafka-сообщения
	if err := json.Unmarshal(msg.Value, &req); err != nil {
		log.Printf("Failed to parse message: %v", err)
		return
	}

	// Проверка наличия токена и текста
	if req.Text == "" || req.Token == "" {
		log.Println("Missing text or token in request. Skipping.")
		return
	}

	// Ищем в Redis по ключу "token:result"
	redisKey := req.Token + ":result"
	cached, err := redis.GetResult(redisKey)
	if err != nil {
		log.Printf("Redis error: %v", err)
		return
	}

	if cached != "" {
		log.Printf("Cache hit. Sending simplified response for token: %s", req.Token)
		err := kafka.PublishSimplifiedResponse(string(msg.Key), cached)
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

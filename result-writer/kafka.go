package main

import (
	"context"
	"encoding/json"
	"log"
	"result-writer/utils"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

func StartKafkaConsumer(cfg Config, redisClient *redis.Client, db *gorm.DB) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.KafkaBroker},
		Topic:   "model_response", // model_response
		GroupID: "model-response-group",
	})
	defer r.Close()

	log.Println("[ModelConsumer] Listening on topic:", cfg.KafkaTopic)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("[ModelConsumer] Read error: %v", err)
			continue
		}

		log.Printf("[ModelConsumer] Received message with Key: %s, Value: %s", string(m.Key), string(m.Value))

		var msg KafkaMessage
		if err := json.Unmarshal(m.Value, &msg); err != nil {
			log.Printf("[ModelConsumer] JSON error: %v", err)
			continue
		}

		log.Printf("[ModelConsumer] Parsed message: original_token=%s, original_text=%s, simplified_text=%s",
			msg.Original.Token, msg.Original.Text, msg.Simplified.Text)

		hash := utils.HashText(msg.Original.Text)

		err = redisClient.Set(ctx, hash, msg.Simplified.Text, 0).Err()
		if err != nil {
			log.Printf("[ModelConsumer] Redis set error: %v", err)
		} else {
			log.Printf("[ModelConsumer] Cached result under hash: %s", hash)
		}

		login, err := GetUserLoginByToken(redisClient, msg.Original.Token)
		if err != nil {
			log.Printf("[ModelConsumer] Redis token error: %v", err)
			continue
		}
		log.Printf("[ModelConsumer] Resolved login: %s", login)

		result := SimplificationResult{
			Original:   msg.Original.Text,
			Simplified: msg.Simplified.Text,
			UserLogin:  login,
		}
		if err := db.Create(&result).Error; err != nil {
			log.Printf("[ModelConsumer] DB insert error: %v", err)
		} else {
			log.Printf("[ModelConsumer] Saved result to DB for user: %s", login)
		}
	}
}

func StartResultRequestConsumer(cfg Config, redisClient *redis.Client, db *gorm.DB) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.KafkaBroker},
		Topic:   "result_request",
		GroupID: "result-request-group",
	})
	defer r.Close()

	writer := &kafka.Writer{
		Addr:     kafka.TCP(cfg.KafkaBroker),
		Topic:    "result_response",
		Balancer: &kafka.LeastBytes{},
	}

	log.Println("[RequestConsumer] Listening on topic: result_request")

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("[RequestConsumer] Read error: %v", err)
			continue
		}

		log.Printf("[RequestConsumer] Received message with Key: %s, Value: %s", string(msg.Key), string(msg.Value))

		correlationID := string(msg.Key)
		if correlationID == "" {
			log.Println("[RequestConsumer] Empty message key (correlationID), skipping")
			continue
		}

		var req ResultRequest
		if err := json.Unmarshal(msg.Value, &req); err != nil {
			log.Printf("[RequestConsumer] JSON error: %v", err)
			continue
		}

		log.Printf("[RequestConsumer] Parsed request: token=%s, correlation_id=%s", req.Token, correlationID)

		req.CorrelationID = correlationID

		login, err := GetUserLoginByToken(redisClient, req.Token)
		if err != nil {
			log.Printf("[RequestConsumer] Redis token error: %v", err)
			continue
		}
		log.Printf("[RequestConsumer] Resolved login: %s", login)

		var results []SimplificationResult
		if err := db.Where("user_login = ?", login).Find(&results).Error; err != nil {
			log.Printf("[RequestConsumer] DB query error: %v", err)
			continue
		}
		log.Printf("[RequestConsumer] Fetched %d results from DB for user: %s", len(results), login)

		response := ResultResponse{
			UserLogin: login,
			Results:   make([]SimplificationPayload, 0, len(results)),
		}
		for _, r := range results {
			response.Results = append(response.Results, SimplificationPayload{
				Original:   r.Original,
				Simplified: r.Simplified,
				CreatedAt:  r.CreatedAt,
			})
		}

		respBytes, _ := json.Marshal(response)
		log.Printf("[RequestConsumer] Sending response with correlation_id=%s, payload_len=%d", correlationID, len(respBytes))

		err = writer.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(correlationID),
			Value: respBytes,
		})
		if err != nil {
			log.Printf("[RequestConsumer] Kafka write error: %v", err)
		} else {
			log.Printf("[RequestConsumer] Successfully sent response for correlation_id=%s", correlationID)
		}
	}
}

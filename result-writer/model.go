package main

import "time"

type KafkaMessage struct {
	Original   TextData `json:"original"`
	Simplified TextData `json:"simplified"`
}

type TextData struct {
	Text  string `json:"text"`
	Token string `json:"token"`
}

type SimplificationResult struct {
	ID         uint `gorm:"primaryKey"`
	Original   string
	Simplified string
	UserLogin  string
	CreatedAt  time.Time
}

type ResultRequest struct {
	CorrelationID string `json:"-"`
	Token         string `json:"token"`
}

type ResultResponse struct {
	UserLogin string                  `json:"user_login"`
	Results   []SimplificationPayload `json:"results"`
}

type SimplificationPayload struct {
	Original   string    `json:"original"`
	Simplified string    `json:"simplified"`
	CreatedAt  time.Time `json:"created_at"`
}

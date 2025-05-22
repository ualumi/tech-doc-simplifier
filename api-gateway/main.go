package main

import (
	"log"
	"net/http"

	"api-gateway/config"
	"api-gateway/handlers"
	"api-gateway/kafka"
	"api-gateway/middleware"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	kafka.SetupKafka()

	kafkaBroker := config.GetEnv("KAFKA_BROKER", "broker:9092")
	kafkaTopic := config.GetEnv("KAFKA_TOPIC", "user_requests")
	port := config.GetEnv("PORT", "8080")

	kafka.InitConsumer(kafkaBroker, "user_responses")
	defer func() {
		if err := kafka.CloseConsumer(); err != nil {
			log.Println("Ошибка при закрытии Kafka consumer:", err)
		}
	}()

	kafka.InitProducer(kafkaBroker, kafkaTopic)
	defer func() {
		if err := kafka.CloseProducer(); err != nil {
			log.Println("Ошибка при закрытии Kafka producer:", err)
		}
	}()

	//for /history
	kafka.InitResultResponseConsumer(kafkaBroker)

	kafka.InitResultRequestProducer(kafkaBroker)
	defer func() {
		if err := kafka.CloseResultRequestProducer(); err != nil {
			log.Printf("Failed to close result request producer: %v", err)
		}
	}()

	r := mux.NewRouter()

	r.Handle("/simplify", middleware.AuthMiddleware(http.HandlerFunc(handlers.SimplifyHandler))).Methods("POST", "OPTIONS")
	r.Handle("/history", middleware.AuthMiddleware(http.HandlerFunc(handlers.HistoryHandler))).Methods("GET", "OPTIONS")

	wrappedRouter := middleware.CorrelationIDMiddleware(r)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://87.228.89.190"},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(wrappedRouter)

	log.Println("API Gateway запущен на порту:", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal("Ошибка сервера:", err)
	}
}

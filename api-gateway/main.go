package main

import (
	"log"
	"net/http"

	"api-gateway/config"
	"api-gateway/handlers"
	"api-gateway/kafka"
	"api-gateway/middleware"

	"github.com/gorilla/mux"
)

func main() {
	kafka.SetupKafka()
	// Получаем переменные окружения или дефолтные значения
	kafkaBroker := config.GetEnv("KAFKA_BROKER", "broker:9092")
	kafkaTopic := config.GetEnv("KAFKA_TOPIC", "user_requests")
	port := config.GetEnv("PORT", "8080")

	// Инициализируем Kafka-писатель
	kafka.InitProducer(kafkaBroker, kafkaTopic)
	defer func() {
		if err := kafka.CloseProducer(); err != nil {
			log.Println("Ошибка при закрытии Kafka:", err)
		}
	}()

	// Настройка роутера
	r := mux.NewRouter()
	r.Handle("/simplify", middleware.AuthMiddleware(http.HandlerFunc(handlers.SimplifyHandler))).Methods("POST")

	log.Println("API Gateway запущен на порту:", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Ошибка сервера:", err)
	}
}

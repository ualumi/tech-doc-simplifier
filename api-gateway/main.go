/*
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

		// Инициализируем Kafka-консьюмер
		kafka.InitConsumer(kafkaBroker, "user_responses") // <-- ЭТО ДОБАВЛЕНО

		defer func() {
			if err := kafka.CloseProducer(); err != nil {
				log.Println("Ошибка при закрытии Kafka:", err)
			}
		}()

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
*/
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

	kafkaBroker := config.GetEnv("KAFKA_BROKER", "broker:9092")
	kafkaTopic := config.GetEnv("KAFKA_TOPIC", "user_requests")
	port := config.GetEnv("PORT", "8080")

	// Инициализируем Kafka-консьюмер ДО отправки сообщений
	kafka.InitConsumer(kafkaBroker, "user_responses")
	defer func() {
		if err := kafka.CloseConsumer(); err != nil {
			log.Println("Ошибка при закрытии Kafka consumer:", err)
		}
	}()

	// Инициализируем Kafka-писатель
	kafka.InitProducer(kafkaBroker, kafkaTopic)
	defer func() {
		if err := kafka.CloseProducer(); err != nil {
			log.Println("Ошибка при закрытии Kafka producer:", err)
		}
	}()

	// Запускаем отдельный Kafka-консьюмер для model_response
	kafka.InitModelResponseConsumer(kafkaBroker, "model_responses")
	kafka.StartModelResponseConsumer(func(msg string) {
		log.Println("Ответ от model-service:", msg)
	})
	defer func() {
		if err := kafka.CloseModelResponseConsumer(); err != nil {
			log.Println("Ошибка при закрытии model_responses consumer:", err)
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

package main

import (
	"auth-service/db"
	"auth-service/handlers"
	"auth-service/kafka"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Ошибка загрузки .env файла:", err)
	}
}

func main() {
	db.InitPostgres()
	db.InitRedis()
	go kafka.StartConsumer()

	http.HandleFunc("/registration", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler) // добавили

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Println("Auth service running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

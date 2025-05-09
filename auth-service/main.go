package main

import (
	"auth-service/db"
	"auth-service/handlers"
	"auth-service/kafka"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
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

	mux := http.NewServeMux()
	mux.HandleFunc("/registration", handlers.RegisterHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true, // 👈 включи логирование CORS
	}).Handler(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Println("Auth service running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}

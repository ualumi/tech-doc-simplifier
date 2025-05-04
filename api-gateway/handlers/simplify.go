package handlers

import (
	"api-gateway/kafka"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type SimplifyRequest struct {
	Text string `json:"text"`
}

// SimplifyHandler обработчик для упрощения текста
func SimplifyHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем токен из заголовка Authorization
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	// Убираем префикс "Bearer " из токена (если он есть)
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	// Читаем тело запроса с текстом
	var req SimplifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Генерируем уникальный корреляционный ID
	corrID := uuid.New().String()

	// Публикуем запрос в Kafka, передаем текст и токен
	err := kafka.PublishSimplifyRequest(corrID, req.Text, token)
	if err != nil {
		http.Error(w, "Failed to send request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Чтение ответа из Kafka
	response, err := kafka.ReadResponse(corrID)
	if err != nil {
		http.Error(w, "Failed to read response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

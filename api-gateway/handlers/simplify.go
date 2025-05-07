package handlers

import (
	"api-gateway/config"
	"api-gateway/kafka"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type SimplifyRequest struct {
	Text string `json:"text"`
}

type SimplifyFullResponse struct {
	UserResponse  string `json:"user_response"`
	ModelResponse string `json:"model_response"`
}

func SimplifyHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	var req SimplifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	corrID := uuid.New().String()

	err := kafka.PublishSimplifyRequest(corrID, req.Text, token)
	if err != nil {
		http.Error(w, "Failed to send request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	userRaw, err := kafka.ReadResponse(corrID)
	if err != nil {
		http.Error(w, "Failed to read user response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Извлекаем фактический user response
	userParts := strings.SplitN(userRaw, ":", 2)
	userResp := strings.TrimSpace(userParts[len(userParts)-1]) // на случай, если нет ':'

	// Если Unauthorized — сразу ответ
	if userResp == "Unauthorized" {
		resp := SimplifyFullResponse{
			UserResponse:  userResp,
			ModelResponse: "",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(resp)
		return
	}

	modelResp, err := kafka.ReadModelResponseOnce(
		config.GetEnv("KAFKA_BROKER", "broker:9092"),
		"model_response",
		corrID,
	)
	if err != nil {
		http.Error(w, "Failed to read model response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := SimplifyFullResponse{
		UserResponse:  userResp,
		ModelResponse: modelResp,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

package handlers

import (
	"api-gateway/kafka"
	"api-gateway/middleware"
	"net/http"
	"strings"
)

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	token = strings.TrimSpace(token)

	corrID := middleware.GetCorrelationID(r.Context())

	if err := kafka.PublishResultRequest(corrID, token); err != nil {
		http.Error(w, "Failed to send history request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseJSON, err := kafka.ReadResultResponse(corrID)
	if err != nil {
		http.Error(w, "Failed to read history response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(responseJSON))
}

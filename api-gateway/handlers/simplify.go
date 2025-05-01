package handlers

import (
	"api-gateway/kafka"
	"encoding/json"
	"net/http"
)

type SimplifyRequest struct {
	Text string `json:"text"`
}

func SimplifyHandler(w http.ResponseWriter, r *http.Request) {
	var req SimplifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := kafka.PublishSimplifyRequest(req.Text)
	if err != nil {
		http.Error(w, "Failed to queue request"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Request received"))
}

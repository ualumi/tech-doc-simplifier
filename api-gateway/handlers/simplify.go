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

func SimplifyHandler(w http.ResponseWriter, r *http.Request) {
	var req SimplifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	corrID := uuid.New().String()

	err := kafka.PublishSimplifyRequest(corrID, req.Text)
	if err != nil {
		http.Error(w, "Failed to send request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := kafka.ReadResponse(corrID)
	if err != nil {
		http.Error(w, "Failed to read response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

/*package handlers

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
}*/

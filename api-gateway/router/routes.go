package router

import (
	"api-gateway/handlers"
	"api-gateway/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.AuthMiddleware)
	r.HandleFunc("/simplify", handlers.SimplifyHandler).Methods("POST")
	r.HandleFunc("/history", handlers.HistoryHandler).Methods("GET") // ← добавлено
	return r
}

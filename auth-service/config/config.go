package config

import (
	"os"
)

// GetEnv возвращает значение переменной окружения или fallback, если она не установлена
func GetEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

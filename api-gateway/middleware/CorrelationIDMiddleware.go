package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const correlationIDKey = contextKey("correlationID")

func CorrelationIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		corrID := uuid.New().String()
		ctx := context.WithValue(r.Context(), correlationIDKey, corrID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetCorrelationID(ctx context.Context) string {
	if val, ok := ctx.Value(correlationIDKey).(string); ok {
		return val
	}
	return ""
}

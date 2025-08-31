// Package middleware содержит HTTP middleware для работы с контекстом,
// включая генерацию уникального идентификатора запроса (RequestID).
package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// RequestIDKey уникальная структура для ключа запроса
type RequestIDKey struct{}

// RequestIDMiddleware добавляем RequestID к каждому запросу
func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		ctx := context.WithValue(r.Context(), RequestIDKey{}, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetRequestID получаем RequestID из контекста
func GetRequestID(ctx context.Context) string {
	requestID := ctx.Value(RequestIDKey{}).(string)
	return requestID
}

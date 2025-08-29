// Package logger пакет для настройки logger
package logger

import (
	"context"
	"log/slog"
	"os"
	"strings"

	"github.com/SKharchenko87/foodix/internal/middleware"
	"github.com/SKharchenko87/foodix/pkg/config"
)

// InitLogger инициализация logger
func InitLogger(configLogger config.Logger) *slog.Logger {
	formatLevel := strings.ToLower(configLogger.Level)

	var level slog.Level
	if err := level.UnmarshalText([]byte(formatLevel)); err != nil {
		level = slog.LevelInfo // default значение
	}

	var handler slog.Handler
	opts := &slog.HandlerOptions{
		Level: level,
	}

	switch strings.ToLower(configLogger.Format) {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, opts)
	default:
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	handler = CustomHandlerLogger{handler}

	return slog.New(handler)
}

// CustomHandlerLogger кастомный logger обработчик добавляющий request_id
type CustomHandlerLogger struct {
	handler slog.Handler
}

// Enabled делегирует вызов вложенному обработчику
func (c CustomHandlerLogger) Enabled(ctx context.Context, level slog.Level) bool {
	return c.handler.Enabled(ctx, level)
}

// Handle добавляем из контекста request_id
func (c CustomHandlerLogger) Handle(ctx context.Context, record slog.Record) error {
	newRecord := record
	if requestID, ok := ctx.Value(middleware.RequestIDKey{}).(string); ok {
		requestIDAttr := slog.String("request_id", requestID)
		newRecord.AddAttrs(requestIDAttr)
	}
	return c.handler.Handle(ctx, newRecord)
}

// WithAttrs делегирует вызов вложенному обработчику
func (c CustomHandlerLogger) WithAttrs(attrs []slog.Attr) slog.Handler {
	return c.handler.WithAttrs(attrs)
}

// WithGroup делегирует вызов вложенному обработчику
func (c CustomHandlerLogger) WithGroup(name string) slog.Handler {
	return c.handler.WithGroup(name)
}

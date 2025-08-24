// Package logger пакет для настройки logger
package logger

import (
	"log/slog"
	"os"
	"strings"

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
	opts := &slog.HandlerOptions{Level: level}

	switch strings.ToLower(configLogger.Format) {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, opts)
	default:
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}

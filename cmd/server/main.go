// Package main точка входа
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/SKharchenko87/foodix/internal/application"
	internalconfig "github.com/SKharchenko87/foodix/internal/config"
)

func main() {
	if err := runApp(); err != nil {
		slog.Error("Application failed", "error", err)
		os.Exit(1)
	}
}

func runApp() error {
	// Определяем путь к конфигурации
	configPath, ok := os.LookupEnv("FOODIX_CONFIG")
	if !ok {
		return fmt.Errorf("environment variable FOODIX_CONFIG not found")
	}

	// Создаем контекст, который отслеживает сигналы завершения
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// Применяем конфигурацию
	var cfg = internalconfig.NewYAMLConfig()
	if err := cfg.Load(configPath); err != nil {
		return err
	}

	// Запускаем приложение
	app, err := application.NewApplication(cfg)
	if err != nil {
		return err
	}
	if err = app.Start(ctx); err != nil {
		return err
	}

	// ToDo graceful shutdown.

	return nil
}

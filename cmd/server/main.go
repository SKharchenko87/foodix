// Package main точка входа
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// Создаем приложение
	app, err := application.NewApplication(cfg)
	if err != nil {
		return err
	}

	// Запускаем приложение в отдельной goroutine, что бы отрабатывал graceful shutdown
	go func() {
		if err = app.Start(ctx); err != nil {
			slog.Error("Application failed", "error", err)
			cancel()
		}
	}()

	// Graceful shutdown полное завершение работы
	<-ctx.Done()
	slog.Info("graceful shutting down...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := app.Stop(shutdownCtx); err != nil {
		return fmt.Errorf("graceful shutdown failed: %w", err)
	}

	slog.Info("graceful shutdown complete")
	return nil
}

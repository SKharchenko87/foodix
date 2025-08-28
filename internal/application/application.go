// Package application пакет для управления жизненным циклом приложения
package application

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/SKharchenko87/foodix/internal/repository"
	"github.com/SKharchenko87/foodix/internal/service"
	"github.com/SKharchenko87/foodix/internal/transport/server"
	"github.com/SKharchenko87/foodix/pkg/config"
	"github.com/SKharchenko87/foodix/pkg/logger"
)

// Application todo
type Application struct {
	cfg        config.Config
	server     server.HTTPServer
	repository repository.ProductRepository
	logger     *slog.Logger
}

// NewApplication todo
func NewApplication(cfg config.Config) (*Application, error) {
	appLogger := logger.InitLogger(cfg.GetLogger())
	appLogger.Info("Creating new application")

	// Инициализируем репозиторий
	productRepo, err := repository.NewRepository(context.Background(), cfg.GetRepo())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize repository: %w", err)
	}

	// Инициализируем сервис с бизнес логикой
	productService := service.NewProductService(productRepo, appLogger)

	// Инициализируем сервер
	httpServer := server.NewHTTPServer(cfg.GetServer(), productService, appLogger)

	return &Application{
		cfg:        cfg,
		server:     httpServer,
		repository: productRepo,
		logger:     appLogger,
	}, nil
}

// Start запуск приложения
func (app *Application) Start(ctx context.Context) error {

	// Инициализируем logger
	app.logger.InfoContext(ctx, "Starting application")

	// Запускам сервер
	if err := app.server.RunServer(ctx); err != nil {
		app.logger.ErrorContext(ctx, "Failed to start server", "error", err)
		return err
	}

	return nil
}

// Stop остановка приложения
func (app *Application) Stop(ctx context.Context) error {
	app.repository.Close()
	err := app.server.Shutdown(ctx)
	if err != nil {
		app.logger.ErrorContext(ctx, "Failed to stop server", "error", err)
		return err
	}
	return nil
}

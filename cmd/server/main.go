// Package main точка входа
package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	internalconfig "github.com/SKharchenko87/foodix/internal/config"
	"github.com/SKharchenko87/foodix/internal/repository"
	"github.com/SKharchenko87/foodix/internal/server"
	"github.com/SKharchenko87/foodix/internal/service"
	"github.com/SKharchenko87/foodix/pkg/config"
	"github.com/SKharchenko87/foodix/pkg/logger"
)

type appArgs struct {
	configPath string
}

func main() {
	if err := runApp(); err != nil {
		slog.Error("Application failed", "error", err)
		os.Exit(1)
	}
}

func runApp() error {
	// Создаем контекст, который отслеживает сигналы завершения
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// Получаем аргументы запуска
	args, err := parseArgs()
	if err != nil {
		return err
	}

	// Применяем конфигурацию
	var cfg config.Config = internalconfig.NewYAMLConfig()
	if err = cfg.Load(args.configPath); err != nil {
		return err
	}

	// Инициализируем logger
	log := logger.InitLogger(cfg.GetLogger())
	log.InfoContext(ctx, "Starting application", "config_path", args.configPath)

	// Настраиваем репозиторий
	productRepo, err := repository.NewRepository(ctx, cfg.GetRepo())
	if err != nil {
		log.ErrorContext(ctx, "Failed to initialize repository", "error", err)
		return err
	}
	if productRepo != nil {
		defer productRepo.Close()
	}

	// Сервис с бизнес логикой
	productService := service.NewProductService(productRepo)

	// Запускам сервер
	httpServer := server.NewHTTPServer(cfg.GetServer(), productService, log)
	if err = httpServer.RunServer(ctx); err != nil {
		log.ErrorContext(ctx, "Failed to start server", "error", err)
	}

	// ToDo graceful shutdown.

	log.InfoContext(ctx, "Application stopped successfully")
	return nil
}

func parseArgs() (*appArgs, error) {
	var pathConfig string // Получаем путь к файлу конфигурации
	flgs := flag.NewFlagSet("server", flag.ExitOnError)
	flgs.StringVar(&pathConfig, "config", "", "path to config file")

	if err := flgs.Parse(os.Args[1:]); err != nil {
		return nil, err
	}

	if pathConfig == "" {
		return nil, flag.ErrHelp
	}

	return &appArgs{pathConfig}, nil
}

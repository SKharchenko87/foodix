// Package main точка входа
package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/SKharchenko87/foodix/internal/application"
	internalconfig "github.com/SKharchenko87/foodix/internal/config"
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
	var cfg = internalconfig.NewYAMLConfig()
	if err = cfg.Load(args.configPath); err != nil {
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

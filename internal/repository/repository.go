// Package repository пакет для интерфейса репозитория
package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/SKharchenko87/foodix/internal/models"
	"github.com/SKharchenko87/foodix/internal/repository/postgres"
	"github.com/SKharchenko87/foodix/pkg/config"
)

// ProductRepository интерфейс для подключения источников данных
type ProductRepository interface {
	Close()
	GetProduct(ctx context.Context, name string) (*models.Product, error)
}

// NewRepository фабрика создает экземпляры разных реализаций ProductRepository
func NewRepository(ctx context.Context, cfg config.Repo) (ProductRepository, error) {
	switch strings.ToLower(cfg.Name) {
	case "postgres":
		return postgres.NewPostgres(ctx, cfg)
	default:
		return nil, fmt.Errorf("repo not implemented")
	}
}

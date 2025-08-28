// Package repository пакет для интерфейса репозитория
package repository

import (
	"context"

	"github.com/SKharchenko87/foodix/internal/domain/models"
)

// ProductRepository интерфейс для подключения источников данных
type ProductRepository interface {
	Close()
	GetProduct(ctx context.Context, name string) (*models.Product, error)
}

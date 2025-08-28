// Package service пакет для бизнес логики
package service

import (
	"context"

	"github.com/SKharchenko87/foodix/internal/domain/models"
)

// ProductService структура для сервиса продуктов
type ProductService interface {
	GetProduct(ctx context.Context, name string) (*models.Product, error)
}

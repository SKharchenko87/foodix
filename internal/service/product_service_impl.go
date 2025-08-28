// Package service пакет для бизнес логики
package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/SKharchenko87/foodix/internal/domain/models"
	"github.com/SKharchenko87/foodix/internal/repository"
)

// ProductServiceImpl структура для сервиса продуктов
type ProductServiceImpl struct {
	repo   repository.ProductRepository
	logger *slog.Logger
}

// NewProductService возвращает экземпляр структуры сервиса продуктов
func NewProductService(repo repository.ProductRepository, logger *slog.Logger) ProductService {
	res := ProductServiceImpl{repo: repo, logger: logger}
	return &res
}

// GetProduct получаем продукт и проводим над ним бизнес действия
func (p *ProductServiceImpl) GetProduct(ctx context.Context, name string) (*models.Product, error) {
	product, err := p.repo.GetProduct(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("product %s error %w", name, err)
	}
	return product, nil
}

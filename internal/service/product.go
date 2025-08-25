// Package service пакет для бизнес логики
package service

import (
	"context"
	"fmt"

	"github.com/SKharchenko87/foodix/internal/models"
	"github.com/SKharchenko87/foodix/internal/repository"
)

// ProductService структура для сервиса продуктов
type ProductService struct {
	repo repository.ProductRepository
}

// NewProductService возвращает экземпляр структуры сервиса продуктов
func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// GetProduct получаем продукт и проводим над ним бизнес действия
func (p *ProductService) GetProduct(ctx context.Context, name string) (*models.Product, error) {
	product, err := p.repo.GetProduct(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("product %s error %w", name, err)
	}
	return product, nil
}

// Package postgres пакет для подключения к БД postgres
package postgres

import (
	"context"

	"github.com/SKharchenko87/foodix/internal/domain/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresRepository структура для подключения к БД
type Repository struct {
	pool *pgxpool.Pool
}

// Close закрываем пул соединений
func (p Repository) Close() {
	p.pool.Close()
}

// GetProduct получаем продукт из БД postgres
func (p Repository) GetProduct(ctx context.Context, name string) (*models.Product, error) {
	r := p.pool.QueryRow(
		ctx,
		`select 
    			name, 
    			protein, 
    			fat, 
    			carbohydrate, 
    			kcal 
			from 
			    public.product p 
			where 
			    name = $1`,
		name,
	)
	res := &models.Product{}
	err := r.Scan(&res.Name, &res.Protein, &res.Fat, &res.Carbohydrate, &res.Kcal)
	if err != nil {
		return nil, err
	}
	return res, nil
}

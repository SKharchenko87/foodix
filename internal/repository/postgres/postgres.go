// Package postgres пакет для подключения к БД postgres
package postgres

import (
	"context"
	"fmt"

	"github.com/SKharchenko87/foodix/internal/models"
	"github.com/SKharchenko87/foodix/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Postgres структура для подключения к БД
type Postgres struct {
	pool *pgxpool.Pool
}

// NewPostgres возвращает экземпляр для подключения к БД postgres
func NewPostgres(ctx context.Context, cfg config.Repo) (*Postgres, error) {
	pool, err := pgxpool.New(ctx, cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("could not connect pool to postgres: %w", err)
	}
	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not ping postgres: %w", err)
	}
	pg := &Postgres{pool: pool}
	return pg, nil
}

// Close закрываем пул соединений
func (p Postgres) Close() {
	p.pool.Close()
}

// GetProduct получаем продукт из БД postgres
func (p Postgres) GetProduct(ctx context.Context, name string) (*models.Product, error) {
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

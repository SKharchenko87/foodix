package postgres

import (
	"context"
	"fmt"

	"github.com/SKharchenko87/foodix/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPostgresRepository возвращает экземпляр для подключения к БД postgres
func NewPostgresRepository(ctx context.Context, cfg config.Repo) (*Repository, error) {
	if cfg.Name == "" {
		return nil, fmt.Errorf("repository name not specified")
	}

	pool, err := pgxpool.New(ctx, cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("could not connect pool to postgres: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not ping postgres: %w", err)
	}

	return &Repository{pool: pool}, nil
}

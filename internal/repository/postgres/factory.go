package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/SKharchenko87/foodix/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPostgresRepository возвращает экземпляр для подключения к БД postgres
func NewPostgresRepository(ctx context.Context, cfg config.Repo, logger *slog.Logger) (*Repository, error) {
	if cfg.Name == "" {
		return nil, fmt.Errorf("repository name not specified")
	}
	logger.Info("Connecting to postgresql...")
	pool, err := pgxpool.New(ctx, cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("could not connect pool to postgres: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not ping postgres: %w", err)
	}

	logger.Info("Connected to postgresql")

	return &Repository{pool: pool, logger: logger}, nil
}

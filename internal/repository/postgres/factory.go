package postgres

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewPostgresRepository возвращает экземпляр для подключения к БД postgres
func NewPostgresRepository(ctx context.Context, logger *slog.Logger) (*Repository, error) {
	logger.Info("Connecting to postgresql...")
	connStr, err := generateConnectPostgresString()
	if err != nil {
		return nil, fmt.Errorf("could not generate connection string to postgres: %w", err)
	}
	pool, err := pgxpool.New(ctx, connStr)
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

func generateConnectPostgresString() (string, error) {
	var err error
	user, exists := os.LookupEnv("POSTGRES_USER")
	if !exists {
		err = errors.Join(err, errors.New("POSTGRES_USER not found"))
	}
	password, exists := os.LookupEnv("POSTGRES_PASSWORD")
	if !exists {
		err = errors.Join(err, errors.New("POSTGRES_PASSWORD not found"))
	}
	db, exists := os.LookupEnv("POSTGRES_DB")
	if !exists {
		err = errors.Join(err, errors.New("POSTGRES_DB not found"))
	}
	host, exists := os.LookupEnv("POSTGRES_HOST")
	if !exists {
		err = errors.Join(err, errors.New("POSTGRES_HOST not found"))
	}
	portStr, exists := os.LookupEnv("POSTGRES_PORT")
	if !exists {
		err = errors.Join(err, errors.New("POSTGRES_PORT not found"))
	}

	port, tmpErr := strconv.Atoi(portStr)
	if tmpErr != nil {
		err = errors.Join(err, errors.New("POSTGRES_PORT must be an integer"), tmpErr)
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, db), err
}

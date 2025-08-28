package repository

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/SKharchenko87/foodix/internal/repository/postgres"
	"github.com/SKharchenko87/foodix/pkg/config"
)

// NewRepository фабрика создает экземпляры разных реализаций ProductRepository
func NewRepository(ctx context.Context, cfg config.Repo, logger *slog.Logger) (ProductRepository, error) {
	repoType := strings.ToLower(cfg.Name)
	switch repoType {
	case "postgres":
		return postgres.NewPostgresRepository(ctx, cfg, logger)
	default:
		return nil, fmt.Errorf("repository type %s not supported", repoType)
	}
}

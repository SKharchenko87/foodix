// Package server пакет для http сервера
package server

import (
	"context"
)

// HTTPServer интерфейс для http сервера
type HTTPServer interface {
	RunServer(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

// Package server пакет для http сервера
package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/SKharchenko87/foodix/internal/service"
	"github.com/SKharchenko87/foodix/internal/transport/handler/product"
	"github.com/SKharchenko87/foodix/pkg/config"
)

// NewHTTPServer возвращает экземпляр HTTPServer
func NewHTTPServer(cfg config.Server, productService service.ProductService, logger *slog.Logger) HTTPServer {
	mux := http.NewServeMux()
	GetProductByNameHandler := product.NewGetProductByNameHandler(productService, logger)
	mux.HandleFunc("GET /product", GetProductByNameHandler.Handle)

	addr := cfg.Host + ":" + strconv.Itoa(cfg.Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second, // ToDo вынести в config
		WriteTimeout: 5 * time.Second, // ToDo вынести в config
		IdleTimeout:  120 * time.Second,
	}
	res := HTTPServerImpl{server, logger}
	return &res
}

// HTTPServerImpl структура для http сервера
type HTTPServerImpl struct {
	srv    *http.Server
	logger *slog.Logger
}

// RunServer запуск веб сервера
func (h *HTTPServerImpl) RunServer(ctx context.Context) error {
	h.logger.InfoContext(ctx, "Сервер запущен", "addr", h.srv.Addr)
	if err := h.srv.ListenAndServe(); err != nil {
		return fmt.Errorf("error starting server %w", err)
	}
	return nil
}

// Shutdown выключение сервера
func (h *HTTPServerImpl) Shutdown(ctx context.Context) error {
	h.logger.InfoContext(ctx, "shutting down http server")
	return h.srv.Shutdown(ctx)
}

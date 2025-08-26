// Package server пакет для http сервера
package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/SKharchenko87/foodix/internal/handler/product"
	"github.com/SKharchenko87/foodix/internal/service"
	"github.com/SKharchenko87/foodix/pkg/config"
)

// HTTPServer структура для http сервера
type HTTPServer struct {
	srv *http.Server
	log *slog.Logger
}

// NewHTTPServer возвращает экземпляр HTTPServer
func NewHTTPServer(cfg config.Server, productService *service.ProductService, log *slog.Logger) *HTTPServer {
	mux := http.NewServeMux()

	GetProductByNameHandler := product.NewGetProductByNameHandler(productService, log)
	mux.HandleFunc("GET /product", GetProductByNameHandler.Handle)

	addr := cfg.Host + ":" + strconv.Itoa(cfg.Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second, // ToDo вынести в config
		WriteTimeout: 5 * time.Second, // ToDo вынести в config
		IdleTimeout:  120 * time.Second,
	}
	return &HTTPServer{server, log}
}

// RunServer запуск веб сервера
func (h *HTTPServer) RunServer(ctx context.Context) error {
	h.log.InfoContext(ctx, "Сервер запущен", "addr", h.srv.Addr)
	if err := h.srv.ListenAndServe(); err != nil {
		return fmt.Errorf("error starting server %w", err)
	}
	return nil
}

// Shutdown выключение сервера
func (h *HTTPServer) Shutdown(ctx context.Context) error {
	h.log.InfoContext(ctx, "shutting down http server")
	return h.srv.Shutdown(ctx)
}

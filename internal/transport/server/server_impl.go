// Package server пакет для http сервера
package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/SKharchenko87/foodix/internal/service"
	"github.com/SKharchenko87/foodix/internal/transport/handler/product"
	"github.com/SKharchenko87/foodix/pkg/config"
)

// Значения по умолчанию для сервера
const (
	defaultReadTimeout  = 5 * time.Second
	defaultWriteTimeout = 5 * time.Second
	defaultIdleTimeout  = 120 * time.Second
)

// NewHTTPServer возвращает экземпляр HTTPServer
func NewHTTPServer(cfg config.Server, productService service.ProductService, logger *slog.Logger) HTTPServer {
	mux := http.NewServeMux()
	GetProductByNameHandler := product.NewGetProductByNameHandler(productService, logger)
	mux.HandleFunc("GET /product", GetProductByNameHandler.Handle)

	addr := cfg.Host + ":" + strconv.Itoa(cfg.Port)

	readTimeout, err := time.ParseDuration(cfg.ReadTimeout)
	if err != nil {
		logger.Warn("Failed to parse ReadTimeout value:", "warning", err, "default", defaultReadTimeout)
		readTimeout = defaultReadTimeout
	}

	writeTimeout, err := time.ParseDuration(cfg.WriteTimeout)
	if err != nil {
		logger.Warn("Failed to parse WriteTimeout value:", "warning", err, "default", defaultWriteTimeout)
		writeTimeout = defaultWriteTimeout
	}

	idleTimeout, err := time.ParseDuration(cfg.IdleTimeout)
	if err != nil {
		logger.Warn("Failed to parse IdleTimeout value:", "warning", err, "default", defaultIdleTimeout)
		idleTimeout = defaultIdleTimeout
	}

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
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
	h.logger.InfoContext(ctx, "Server starting", "addr", h.srv.Addr)
	if err := h.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("error starting server %w", err)
	}
	return nil
}

// Shutdown выключение сервера
func (h *HTTPServerImpl) Shutdown(ctx context.Context) error {
	h.logger.InfoContext(ctx, "Shutting down http server")
	return h.srv.Shutdown(ctx)
}

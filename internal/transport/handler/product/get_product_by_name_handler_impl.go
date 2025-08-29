// Package product пакет для ручек продуктов
package product

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/SKharchenko87/foodix/internal/service"
)

// GetProductByNameHandlerImpl структура для ручки получения продукта
type GetProductByNameHandlerImpl struct {
	productService service.ProductService
	logger         *slog.Logger
}

// NewGetProductByNameHandler возвращает экземпляр для ручки получения продукта
func NewGetProductByNameHandler(service service.ProductService, logger *slog.Logger) GetProductByNameHandler {
	res := GetProductByNameHandlerImpl{service, logger}
	return &res
}

// Handle ручка
func (handler *GetProductByNameHandlerImpl) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	w.Header().Set("Content-Type", "application/json")

	name := r.URL.Query().Get("name")
	if name == "" {
		handler.writeJSONError(w, http.StatusBadRequest, "name is required", handler.logger)
		return
	}

	handler.logger.DebugContext(ctx, "query find product", "name", name)

	product, err := handler.productService.GetProduct(ctx, name)
	if err != nil {
		handler.writeJSONError(w, http.StatusBadRequest, "product not found", handler.logger)
		return
	}

	if err = json.NewEncoder(w).Encode(product); err != nil {
		handler.logger.ErrorContext(ctx, "product json encode error", "name", name, "err", err)
	}

}
func (handler *GetProductByNameHandlerImpl) writeJSONError(w http.ResponseWriter, status int, message string, log *slog.Logger) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(map[string]string{"error": message}); err != nil {
		log.Error("failed to write JSON error response",
			"status", status,
			"message", message,
			"error", err,
		)
	}
}

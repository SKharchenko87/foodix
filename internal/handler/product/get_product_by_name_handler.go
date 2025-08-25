// Package product пакет для ручек продуктов
package product

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/SKharchenko87/foodix/internal/service"
)

// GetProductByNameHandler структура для ручки получения продукта
type GetProductByNameHandler struct {
	productService *service.ProductService
	log            *slog.Logger
}

// NewGetProductByNameHandler возвращает экземпляр для ручки получения продукта
func NewGetProductByNameHandler(service *service.ProductService, log *slog.Logger) *GetProductByNameHandler {
	return &GetProductByNameHandler{service, log}
}

// Handle ручка
func (handler *GetProductByNameHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	w.Header().Set("Content-Type", "application/json")

	name := r.URL.Query().Get("name")
	if name == "" {
		handler.writeJSONError(w, http.StatusBadRequest, "name is required")
		return
	}

	handler.log.Debug("query find product", "name", name)

	product, err := handler.productService.GetProduct(ctx, name)
	if err != nil {
		handler.writeJSONError(w, http.StatusBadRequest, "product not found")
		return
	}

	if err = json.NewEncoder(w).Encode(product); err != nil {
		handler.log.Error("product json encode error", "name", name, "err", err)
	}

}
func (handler *GetProductByNameHandler) writeJSONError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(map[string]string{"error": message}); err != nil {
		handler.log.Error("failed to write JSON error response",
			"status", status,
			"message", message,
			"error", err,
		)
	}
}

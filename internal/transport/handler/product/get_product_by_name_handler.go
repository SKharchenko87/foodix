// Package product пакет для ручек продуктов
package product

import (
	"net/http"
)

// GetProductByNameHandler интерфейс для ручки продукта
type GetProductByNameHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

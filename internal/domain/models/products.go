// Package models пакет бизнес сущностей
package models

// Product структура продукта
type Product struct {
	Name         string
	Protein      float64
	Fat          float64
	Carbohydrate float64
	Kcal         int
}

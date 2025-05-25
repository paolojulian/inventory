package factory

import (
	"paolojulian.dev/inventory/domain/product"
	"paolojulian.dev/inventory/pkg/id"
)

// NewTestProduct returns a product instance with default test values
func NewTestProduct() *product.Product {
	return &product.Product{
		ID:          id.NewUUID(),
		SKU:         "TESTSKU123",
		Name:        "Sample Product",
		Description: "This is a test product.",
		Price:       product.Money{Cents: 4999},
		IsActive:    true,
	}
}

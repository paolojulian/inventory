package product

import "paolojulian.dev/inventory/pkg/id"

type Product struct {
	ID          string
	SKU         SKU
	Name        string
	Description Description
	Price       Money
	IsActive    bool
}

type ProductSummary struct {
	ID    string
	SKU   SKU
	Name  string
	Price Money
}

func NewProduct(sku SKU, name string, description Description, priceCents int) *Product {
	return &Product{
		ID:          id.NewUUID(),
		SKU:         sku,
		Name:        name,
		Description: description,
		Price:       Money{Cents: priceCents},
		IsActive:    true,
	}
}

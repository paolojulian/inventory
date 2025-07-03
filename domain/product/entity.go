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

type ProductFilter struct {
	SearchText *string
	IsActive   *bool
}

type ProductSort struct {
	Field string
	Order string
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

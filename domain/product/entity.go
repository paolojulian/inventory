package product

import (
	"paolojulian.dev/inventory/pkg/id"
	paginationShared "paolojulian.dev/inventory/shared/pagination"
)

type Product struct {
	ID          string      `json:"id"`
	SKU         SKU         `json:"sku"`
	Name        string      `json:"name"`
	Description Description `json:"description"`
	Price       Money       `json:"price"`
	IsActive    bool        `json:"is_active"`
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

// ================================
// SORTING
// ================================
type ProductSortField string

const (
	ProductSortFieldName  ProductSortField = "name"
	ProductSortFieldSKU   ProductSortField = "sku"
	ProductSortFieldPrice ProductSortField = "price"
)

func (f ProductSortField) IsValid() bool {
	return f == ProductSortFieldName || f == ProductSortFieldSKU || f == ProductSortFieldPrice
}

type ProductSort struct {
	Field *ProductSortField           `json:"field,omitempty"`
	Order *paginationShared.SortOrder `json:"order,omitempty"`
}

// ================================
// END SORTING
// ================================

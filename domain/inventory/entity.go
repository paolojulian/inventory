package inventory

import (
	"time"

	"paolojulian.dev/inventory/domain/product"
	"paolojulian.dev/inventory/domain/warehouse"
)

// InventoryItem represents the current stock level of a product in a warehouse
type InventoryItem struct {
	Product     *product.Product
	Warehouse   *warehouse.Warehouse
	Stock       int
	LastUpdated time.Time
}

// StockLevel represents a stock level entry
type StockLevel struct {
	ProductID   string
	WarehouseID string
	Quantity    int
	UpdatedAt   time.Time
}

// InventorySummary represents the overall inventory summary
type InventorySummary struct {
	TotalProducts   int
	TotalStockValue int64 // in cents
	LowStockItems   int
	OutOfStockItems int
	LastUpdated     time.Time
}

// NewInventoryItem creates a new InventoryItem instance
func NewInventoryItem(product *product.Product, warehouse *warehouse.Warehouse, stock int) *InventoryItem {
	return &InventoryItem{
		Product:     product,
		Warehouse:   warehouse,
		Stock:       stock,
		LastUpdated: time.Now(),
	}
}

// IsLowStock checks if the product is low on stock (less than 10 units)
func (ii *InventoryItem) IsLowStock() bool {
	return ii.Stock < 10
}

// IsOutOfStock checks if the product is out of stock
func (ii *InventoryItem) IsOutOfStock() bool {
	return ii.Stock <= 0
}

// GetStockStatus returns the stock status as a string
func (ii *InventoryItem) GetStockStatus() string {
	if ii.IsOutOfStock() {
		return "out_of_stock"
	}
	if ii.IsLowStock() {
		return "low_stock"
	}
	return "in_stock"
}

package factory

import (
	"time"

	"paolojulian.dev/inventory/domain/stock"
	"paolojulian.dev/inventory/pkg/id"
)

func NewTestStockEntry() *stock.StockEntry {
	return &stock.StockEntry{
		ID:            id.NewUUID(),
		QuantityDelta: 50,
		Reason:        stock.ReasonSale,
		CreatedAt:     time.Now(),
		ProductID:     id.NewUUID(),
		WarehouseID:   id.NewUUID(),
		UserID:        id.NewUUID(),
	}
}

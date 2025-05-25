package stock

import (
	"time"
)

type StockEntry struct {
	ID            string
	QuantityDelta int
	Reason        StockReason
	CreatedAt     time.Time
	ProductID     string
	WarehouseID   string
	UserID        string
}

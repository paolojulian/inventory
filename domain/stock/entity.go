package stock

import (
	"fmt"
	"time"

	"paolojulian.dev/inventory/pkg/id"
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

func NewStockEntry(productID, warehouseID, userID string, quantityDelta int, reason StockReason) StockEntry {
	return StockEntry{
		ID:            id.NewUUID(),
		QuantityDelta: quantityDelta,
		Reason:        reason,
		CreatedAt:     time.Now(),
		ProductID:     productID,
		WarehouseID:   warehouseID,
		UserID:        userID,
	}
}

func (s *StockEntry) Validate() error {
	if s.QuantityDelta == 0 {
		return ErrQuantityDeltaZero
	}

	if s.Reason == "" {
		return ErrReasonEmpty
	}

	return nil
}

func (s *StockEntry) String() string {
	return fmt.Sprintf("StockEntry[ID=%s, ProductID=%s, WarehouseID=%s, QuantityDelta=%d, Reason=%s, CreatedAt=%s, UserID=%s]",
		s.ID, s.ProductID, s.WarehouseID, s.QuantityDelta, s.Reason, s.CreatedAt.Format(time.RFC3339), s.UserID)
}

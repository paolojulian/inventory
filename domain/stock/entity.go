package stock

import (
	"fmt"
	"time"

	"paolojulian.dev/inventory/pkg/id"
)

type StockEntry struct {
	ID                 string
	QuantityDelta      int
	Reason             StockReason
	CreatedAt          time.Time
	ProductID          string
	WarehouseID        string
	UserID             string
	SupplierPriceCents *int
	StorePriceCents    *int
	ExpiryDate         *time.Time
	ReorderDate        *time.Time
}

func NewStockEntry(productID, warehouseID, userID string, quantityDelta int, reason StockReason, supplierPriceCents *int, storePriceCents *int, expiryDate *time.Time, reorderDate *time.Time) *StockEntry {
	var resolvedQuantityDelta int = quantityDelta
	switch reason {
	case ReasonDamage,
		ReasonSale,
		ReasonTransferOut:
		resolvedQuantityDelta = resolvedQuantityDelta * -1
	}
	return &StockEntry{
		ID:                 id.NewUUID(),
		QuantityDelta:      resolvedQuantityDelta,
		Reason:             reason,
		CreatedAt:          time.Now(),
		ProductID:          productID,
		WarehouseID:        warehouseID,
		UserID:             userID,
		SupplierPriceCents: supplierPriceCents,
		StorePriceCents:    storePriceCents,
		ExpiryDate:         expiryDate,
		ReorderDate:        reorderDate,
	}
}

func (s *StockEntry) Validate() error {
	if s.QuantityDelta == 0 {
		return ErrQuantityDeltaZero
	}

	if s.Reason == "" {
		return ErrReasonEmpty
	}

	if s.SupplierPriceCents != nil && *s.SupplierPriceCents < 0 {
		return ErrInvalidSupplierPrice
	}
	if s.StorePriceCents != nil && *s.StorePriceCents < 0 {
		return ErrInvalidStorePrice
	}

	return nil
}

func (s *StockEntry) String() string {
	return fmt.Sprintf("StockEntry[ID=%s, ProductID=%s, WarehouseID=%s, QuantityDelta=%d, Reason=%s, CreatedAt=%s, UserID=%s]",
		s.ID, s.ProductID, s.WarehouseID, s.QuantityDelta, s.Reason, s.CreatedAt.Format(time.RFC3339), s.UserID)
}

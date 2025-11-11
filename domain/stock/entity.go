package stock

import (
	"fmt"
	"time"

	"paolojulian.dev/inventory/domain/product"
	"paolojulian.dev/inventory/domain/user"
	"paolojulian.dev/inventory/domain/warehouse"
	"paolojulian.dev/inventory/pkg/id"
)

type StockEntry struct {
	ID                 string      `json:"id"`
	QuantityDelta      int         `json:"quantity_delta"`
	Reason             StockReason `json:"reason"`
	CreatedAt          time.Time   `json:"created_at"`
	ProductID          string      `json:"product_id"`
	WarehouseID        string      `json:"warehouse_id"`
	UserID             string      `json:"user_id"`
	SupplierPriceCents *int        `json:"supplier_price_cents,omitempty"`
	StorePriceCents    *int        `json:"store_price_cents,omitempty"`
	ExpiryDate         *time.Time  `json:"expiry_date,omitempty"`
	ReorderDate        *time.Time  `json:"reorder_date,omitempty"`
}

type StockEntryWithRelations struct {
	StockEntry
	Product   *product.Product     `json:"product,omitempty"`
	Warehouse *warehouse.Warehouse `json:"warehouse,omitempty"`
	User      *user.User           `json:"user,omitempty"`
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

package stock_uc

import (
	"context"
	"time"

	"paolojulian.dev/inventory/domain/stock"
)

type StockEntryInput struct {
	QuantityDelta      int        `json:"quantity_delta"`
	Reason             string     `json:"reason"`
	ProductID          string     `json:"product_id"`
	WarehouseID        string     `json:"warehouse_id"`
	SupplierPriceCents *int       `json:"supplier_price_cents"`
	StorePriceCents    *int       `json:"store_price_cents"`
	ExpiryDate         *time.Time `json:"expiry_date"`
	ReorderDate        *time.Time `json:"reorder_date"`
}

type StockEntryOutput struct {
	StockEntry *stock.StockEntry
}

type StockEntryRepository interface {
	CreateStockEntry(ctx context.Context, stockEntry *stock.StockEntry) (*stock.StockEntry, error)
}

type CreateStockEntryUseCase struct {
	repo StockEntryRepository
}

func NewCreateStockEntryUseCase(repo StockEntryRepository) *CreateStockEntryUseCase {
	return &CreateStockEntryUseCase{repo}
}

func (uc *CreateStockEntryUseCase) Execute(ctx context.Context, input *StockEntryInput, userID string) (*StockEntryOutput, error) {
	if !stock.IsValidStockReason(input.Reason) {
		return nil, ErrInvalidReason
	}

	stockEntry := stock.NewStockEntry(
		input.ProductID,
		input.WarehouseID,
		userID,
		input.QuantityDelta,
		stock.StockReason(input.Reason),
		input.SupplierPriceCents,
		input.StorePriceCents,
		input.ExpiryDate,
		input.ReorderDate,
	)
	if err := stockEntry.Validate(); err != nil {
		return nil, err
	}

	createdStockEntry, err := uc.repo.CreateStockEntry(ctx, &stockEntry)
	if err != nil {
		return nil, err
	}

	return &StockEntryOutput{StockEntry: createdStockEntry}, nil
}

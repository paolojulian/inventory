package stock_uc

import (
	"context"

	"paolojulian.dev/inventory/domain/stock"
)

type StockEntryInput struct {
	QuantityDelta int    `json:"quantity_delta"`
	Reason        string `json:"reason"`
	ProductID     string `json:"product_id"`
	WarehouseID   string `json:"warehouse_id"`
}

type StockEntryOutput struct {
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
	)
	if err := stockEntry.Validate(); err != nil {
		return nil, err
	}

	_, err := uc.repo.CreateStockEntry(ctx, &stockEntry)
	if err != nil {
		return nil, err
	}

	return &StockEntryOutput{}, nil
}

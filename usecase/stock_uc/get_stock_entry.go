package stock_uc

import (
	"context"

	"paolojulian.dev/inventory/domain/stock"
)

type GetStockEntryInput struct {
	StockEntryID string
}

type GetStockEntryOutput struct {
	StockEntry *stock.StockEntry
}

type GetStockEntryRepository interface {
	GetByID(ctx context.Context, stockEntryID string) (stock.StockEntry, error)
}

type GetStockEntryUseCase struct {
	repo GetStockEntryRepository
}

func NewGetStockEntryUseCase(repo GetStockEntryRepository) *GetStockEntryUseCase {
	return &GetStockEntryUseCase{repo}
}

func (uc *GetStockEntryUseCase) Execute(ctx context.Context, stockEntryID string) (stock.StockEntry, error) {
	return uc.repo.GetByID(ctx, stockEntryID)
}

package stock_uc

import (
	"context"

	"paolojulian.dev/inventory/config"
	"paolojulian.dev/inventory/domain/stock"
)

type ListStockEntriesInput struct {
	// filter
	Limit *int `json:"limit,omitempty"`
}

type ListStockEntriesOutput struct {
	StockEntries []*stock.StockEntry
}

type StockRepository interface {
	GetList(ctx context.Context, limit int) ([]*stock.StockEntry, error)
}

type ListStockEntriesUseCase struct {
	repo StockRepository
}

func NewListStockEntriesUseCase(repo StockRepository) *ListStockEntriesUseCase {
	return &ListStockEntriesUseCase{repo}
}

func (uc *ListStockEntriesUseCase) Execute(ctx context.Context, input *ListStockEntriesInput) ([]*stock.StockEntry, error) {
	limit := config.DefaultListLimit
	if input.Limit != nil || *input.Limit <= 0 {
		limit = *input.Limit
	}

	return uc.repo.GetList(ctx, limit)
}

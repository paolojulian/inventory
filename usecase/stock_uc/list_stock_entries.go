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
	Total        int
}

type ListStockEntriesRepository interface {
	GetList(ctx context.Context, limit int) ([]*stock.StockEntry, int, error)
}

type ListStockEntriesUseCase struct {
	repo ListStockEntriesRepository
}

func NewListStockEntriesUseCase(repo ListStockEntriesRepository) *ListStockEntriesUseCase {
	return &ListStockEntriesUseCase{repo}
}

func (uc *ListStockEntriesUseCase) Execute(ctx context.Context, input *ListStockEntriesInput) (*ListStockEntriesOutput, error) {
	limit := config.DefaultListLimit
	if input.Limit != nil || *input.Limit <= 0 {
		limit = *input.Limit
	}

	stockEntries, totalCount, err := uc.repo.GetList(ctx, limit)
	if err != nil {
		return nil, err
	}

	return &ListStockEntriesOutput{
		StockEntries: stockEntries,
		Total:        totalCount,
	}, nil
}

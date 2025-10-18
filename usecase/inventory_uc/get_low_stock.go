package inventory_uc

import (
	"context"

	"paolojulian.dev/inventory/domain/inventory"
)

type GetLowStockInput struct {
	Threshold int `json:"threshold,omitempty"`
}

type GetLowStockOutput struct {
	Stocks []*inventory.InventoryItem `json:"stocks"`
}

type GetLowStockRepository interface {
	GetLowStockProducts(ctx context.Context, threshold int) ([]*inventory.InventoryItem, error)
}

type GetLowStockUseCase struct {
	repo GetLowStockRepository
}

func NewGetLowStockUseCase(repo GetLowStockRepository) *GetLowStockUseCase {
	return &GetLowStockUseCase{repo}
}

func (uc *GetLowStockUseCase) Execute(ctx context.Context, input GetLowStockInput) (*GetLowStockOutput, error) {
	threshold := input.Threshold
	if threshold == 0 {
		threshold = 10 // Default threshold
	}

	stocks, err := uc.repo.GetLowStockProducts(ctx, threshold)
	if err != nil {
		return nil, err
	}

	return &GetLowStockOutput{Stocks: stocks}, nil
}

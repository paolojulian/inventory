package inventory_uc

import (
	"context"

	"paolojulian.dev/inventory/domain/inventory"
)

type GetLowStockInput struct {
	WarehouseID string `json:"warehouse_id" binding:"required"`
	Threshold   int    `json:"threshold,omitempty"`
}

type GetLowStockOutput struct {
	Stocks []*inventory.InventoryItem `json:"stocks"`
}

type GetLowStockRepository interface {
	GetLowStockProducts(ctx context.Context, warehouseID string, threshold int) ([]*inventory.InventoryItem, error)
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

	stocks, err := uc.repo.GetLowStockProducts(ctx, input.WarehouseID, threshold)
	if err != nil {
		return nil, err
	}

	return &GetLowStockOutput{Stocks: stocks}, nil
}

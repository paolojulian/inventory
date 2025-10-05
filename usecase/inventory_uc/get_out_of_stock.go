package inventory_uc

import (
	"context"

	"paolojulian.dev/inventory/domain/inventory"
)

type GetOutOfStockInput struct {
	WarehouseID string `json:"warehouse_id" binding:"required"`
}

type GetOutOfStockOutput struct {
	Stocks []*inventory.InventoryItem `json:"stocks"`
}

type GetOutOfStockRepository interface {
	GetOutOfStockProducts(ctx context.Context, warehouseID string) ([]*inventory.InventoryItem, error)
}

type GetOutOfStockUseCase struct {
	repo GetOutOfStockRepository
}

func NewGetOutOfStockUseCase(repo GetOutOfStockRepository) *GetOutOfStockUseCase {
	return &GetOutOfStockUseCase{repo}
}

func (uc *GetOutOfStockUseCase) Execute(ctx context.Context, input GetOutOfStockInput) (*GetOutOfStockOutput, error) {
	stocks, err := uc.repo.GetOutOfStockProducts(ctx, input.WarehouseID)
	if err != nil {
		return nil, err
	}

	return &GetOutOfStockOutput{Stocks: stocks}, nil
}

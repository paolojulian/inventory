package inventory_uc

import (
	"context"

	"paolojulian.dev/inventory/domain/inventory"
)

type GetAllCurrentStockInput struct {
	WarehouseID string `json:"warehouse_id" binding:"required"`
}

type GetAllCurrentStockOutput struct {
	Stocks []*inventory.InventoryItem `json:"stocks"`
}

type GetAllCurrentStockRepository interface {
	GetAllCurrentStock(ctx context.Context, warehouseID string) ([]*inventory.InventoryItem, error)
}

type GetAllCurrentStockUseCase struct {
	repo GetAllCurrentStockRepository
}

func NewGetAllCurrentStockUseCase(repo GetAllCurrentStockRepository) *GetAllCurrentStockUseCase {
	return &GetAllCurrentStockUseCase{repo}
}

func (uc *GetAllCurrentStockUseCase) Execute(ctx context.Context, input GetAllCurrentStockInput) (*GetAllCurrentStockOutput, error) {
	stocks, err := uc.repo.GetAllCurrentStock(ctx, input.WarehouseID)
	if err != nil {
		return nil, err
	}

	return &GetAllCurrentStockOutput{Stocks: stocks}, nil
}

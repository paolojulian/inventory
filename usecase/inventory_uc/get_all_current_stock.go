package inventory_uc

import (
	"context"

	"paolojulian.dev/inventory/domain/inventory"
)

type GetAllCurrentStockInput struct {
	// No input needed - uses default warehouse
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
	stocks, err := uc.repo.GetAllCurrentStock(ctx, "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		return nil, err
	}

	return &GetAllCurrentStockOutput{Stocks: stocks}, nil
}

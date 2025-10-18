package inventory_uc

import (
	"context"

	"paolojulian.dev/inventory/domain/inventory"
)

type GetOutOfStockInput struct {
	// No input needed - uses default warehouse
}

type GetOutOfStockOutput struct {
	Stocks []*inventory.InventoryItem `json:"stocks"`
}

type GetOutOfStockRepository interface {
	GetOutOfStockProducts(ctx context.Context) ([]*inventory.InventoryItem, error)
}

type GetOutOfStockUseCase struct {
	repo GetOutOfStockRepository
}

func NewGetOutOfStockUseCase(repo GetOutOfStockRepository) *GetOutOfStockUseCase {
	return &GetOutOfStockUseCase{repo}
}

func (uc *GetOutOfStockUseCase) Execute(ctx context.Context, input GetOutOfStockInput) (*GetOutOfStockOutput, error) {
	stocks, err := uc.repo.GetOutOfStockProducts(ctx)
	if err != nil {
		return nil, err
	}

	return &GetOutOfStockOutput{Stocks: stocks}, nil
}

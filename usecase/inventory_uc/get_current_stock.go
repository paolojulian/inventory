package inventory_uc

import (
	"context"

	"paolojulian.dev/inventory/domain/inventory"
)

type GetCurrentStockInput struct {
	ProductID string `json:"product_id" binding:"required"`
}

type GetCurrentStockOutput struct {
	Stock *inventory.InventoryItem `json:"stock"`
}

type GetCurrentStockRepository interface {
	GetCurrentStock(ctx context.Context, productID string) (*inventory.InventoryItem, error)
}

type GetCurrentStockUseCase struct {
	repo GetCurrentStockRepository
}

func NewGetCurrentStockUseCase(repo GetCurrentStockRepository) *GetCurrentStockUseCase {
	return &GetCurrentStockUseCase{repo}
}

func (uc *GetCurrentStockUseCase) Execute(ctx context.Context, input GetCurrentStockInput) (*GetCurrentStockOutput, error) {
	stock, err := uc.repo.GetCurrentStock(ctx, input.ProductID)
	if err != nil {
		return nil, err
	}

	return &GetCurrentStockOutput{Stock: stock}, nil
}

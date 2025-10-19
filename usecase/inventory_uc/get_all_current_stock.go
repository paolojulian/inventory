package inventory_uc

import (
	"context"

	"paolojulian.dev/inventory/domain/inventory"
	paginationShared "paolojulian.dev/inventory/shared/pagination"
)

type GetAllCurrentStockInput struct {
	Pager paginationShared.PagerInput `json:"pager"`
}

type GetAllCurrentStockRepository interface {
	GetAllCurrentStock(ctx context.Context, warehouseID string, pager paginationShared.PagerInput) (*inventory.GetAllStockOutput, error)
}

type GetAllCurrentStockUseCase struct {
	repo GetAllCurrentStockRepository
}

func NewGetAllCurrentStockUseCase(repo GetAllCurrentStockRepository) *GetAllCurrentStockUseCase {
	return &GetAllCurrentStockUseCase{repo}
}

func (uc *GetAllCurrentStockUseCase) Execute(ctx context.Context, input GetAllCurrentStockInput) (*inventory.GetAllStockOutput, error) {
	result, err := uc.repo.GetAllCurrentStock(ctx, "550e8400-e29b-41d4-a716-446655440000", input.Pager)
	if err != nil {
		return nil, err
	}

	return result, nil
}

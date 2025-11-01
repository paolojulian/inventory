package inventory_uc

import (
	"context"

	"paolojulian.dev/inventory/domain/inventory"
	"paolojulian.dev/inventory/domain/warehouse"
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
	result, err := uc.repo.GetAllCurrentStock(ctx, warehouse.DefaultWarehouseID, input.Pager)
	if err != nil {
		return nil, err
	}

	return result, nil
}

package warehouse_uc

import (
	"context"

	"paolojulian.dev/inventory/domain/warehouse"
)

type CreateWarehouseInput struct {
	Address string `json:"address" binding:"required"`
}

type CreateWarehouseOutput struct {
	Warehouse *warehouse.Warehouse
}

type WarehouseRepository interface {
	Create(ctx context.Context, warehouse *warehouse.Warehouse) (*warehouse.Warehouse, error)
}

type CreateWarehouseUseCase struct {
	repo WarehouseRepository
}

func NewCreateWarehouseUseCase(repo WarehouseRepository) *CreateWarehouseUseCase {
	return &CreateWarehouseUseCase{repo}
}

func (uc *CreateWarehouseUseCase) Execute(ctx context.Context, input CreateWarehouseInput) (*CreateWarehouseOutput, error) {
	location := &warehouse.WarehouseLocation{
		Address: input.Address,
	}

	newWarehouse := &warehouse.Warehouse{
		Location: *location,
	}

	createdWarehouse, err := uc.repo.Create(ctx, newWarehouse)
	if err != nil {
		return nil, err
	}

	return &CreateWarehouseOutput{
		Warehouse: createdWarehouse,
	}, nil
}

package warehouse_uc

import (
	"context"

	"paolojulian.dev/inventory/domain/warehouse"
)

type UpdateWarehouseInput struct {
	Address  *string `json:"address,omitempty"`
	IsActive *bool   `json:"is_active,omitempty"`
}

type UpdateWarehouseOutput struct {
	Warehouse *warehouse.Warehouse
}

type UpdateWarehouseRepo interface {
	UpdateByID(ctx context.Context, warehouseID string, warehouse *warehouse.WarehousePatch) (*warehouse.Warehouse, error)
}

type UpdateWarehouseUseCase struct {
	repo UpdateWarehouseRepo
}

func NewUpdateWarehouseUseCase(repo UpdateWarehouseRepo) *UpdateWarehouseUseCase {
	return &UpdateWarehouseUseCase{repo}
}

func (uc *UpdateWarehouseUseCase) Execute(ctx context.Context, warehouseID string, input UpdateWarehouseInput) (*UpdateWarehouseOutput, error) {
	var location *warehouse.WarehouseLocation
	if input.Address != nil && *input.Address != "" {
		location = &warehouse.WarehouseLocation{
			Address: *input.Address,
		}
	}

	warehouseFieldsToUpdate := &warehouse.WarehousePatch{
		Location: location,
		IsActive: input.IsActive,
		// Other fields are unchanged
	}

	newWarehouse, err := uc.repo.UpdateByID(ctx, warehouseID, warehouseFieldsToUpdate)
	if err != nil {
		return nil, err
	}

	return &UpdateWarehouseOutput{Warehouse: newWarehouse}, nil
}

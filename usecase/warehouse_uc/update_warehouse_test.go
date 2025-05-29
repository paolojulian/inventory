package warehouse_uc_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"paolojulian.dev/inventory/config"
	"paolojulian.dev/inventory/domain/warehouse"
	"paolojulian.dev/inventory/tests/factory"
	"paolojulian.dev/inventory/usecase/warehouse_uc"
)

// --- Mock Repo ---

type MockUpdateWarehouseRepo struct {
	updated    *warehouse.Warehouse
	warehouses map[string]*warehouse.Warehouse
}

func (r *MockUpdateWarehouseRepo) UpdateByID(ctx context.Context, warehouseID string, warehouse *warehouse.WarehousePatch) (*warehouse.Warehouse, error) {
	if existingWarehouse, exists := r.warehouses[warehouseID]; exists {
		if warehouse.IsActive != nil {
			existingWarehouse.IsActive = *warehouse.IsActive
		}
		if warehouse.Location != nil {
			existingWarehouse.Location = *warehouse.Location
		}

		return existingWarehouse, nil
	}

	return nil, warehouse_uc.ErrWarehouseNotFound
}

// --- Tests ---

func TestUpdateWarehouse__ValidInput(t *testing.T) {
	mockWarehouse := factory.NewTestWarehouse()

	repo := &MockUpdateWarehouseRepo{
		warehouses: map[string]*warehouse.Warehouse{
			mockWarehouse.ID: mockWarehouse,
		},
	}

	uc := warehouse_uc.NewUpdateWarehouseUseCase(repo)

	var address string = "New Address"
	var is_active bool = false

	input := warehouse_uc.UpdateWarehouseInput{
		Address:  &address,
		IsActive: &is_active,
	}

	result, err := uc.Execute(context.Background(), mockWarehouse.ID, input)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, mockWarehouse.ID, result.Warehouse.ID)
	assert.Equal(t, address, result.Warehouse.Location.Address)
	assert.Equal(t, is_active, result.Warehouse.IsActive)
}

func TestUpdateWarehouse__NotFound(t *testing.T) {
	repo := &MockUpdateWarehouseRepo{
		warehouses: map[string]*warehouse.Warehouse{},
	}

	uc := warehouse_uc.NewUpdateWarehouseUseCase(repo)

	input := warehouse_uc.UpdateWarehouseInput{
		Address:  config.StringPointer("New Address"),
		IsActive: config.BoolPointer(false),
	}

	_, err := uc.Execute(context.Background(), "non-existing-product-id", input)
	if err == nil {
		t.Fatal("expected error but nothing given")
	}
}

package warehouse_uc_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	warehouseDomain "paolojulian.dev/inventory/domain/warehouse"
	"paolojulian.dev/inventory/pkg/id"
	"paolojulian.dev/inventory/tests/middleware"
	"paolojulian.dev/inventory/usecase/warehouse_uc"
)

type MockWarehouseRepository struct {
	warehouses map[string]*warehouseDomain.Warehouse
}

func (repo *MockWarehouseRepository) Create(ctx context.Context, warehouse *warehouseDomain.Warehouse) (*warehouseDomain.Warehouse, error) {
	newWarehouse := &warehouseDomain.Warehouse{
		ID:            id.NewUUID(),
		Location:      warehouse.Location,
		IsActive:      true,
		CreatedBy:     middleware.FakeUserID,
		LastUpdatedBy: middleware.FakeUserID,
	}

	return newWarehouse, nil
}

func TestCreateWarehouse__Success(t *testing.T) {
	repo := &MockWarehouseRepository{}
	uc := warehouse_uc.NewCreateWarehouseUseCase(repo)

	input := warehouse_uc.CreateWarehouseInput{
		Address: "#1321 Biglaan Street, Saint Michael Mayapis",
	}

	result, err := uc.Execute(context.Background(), input)

	assert.NoError(t, err)
	assert.NotEmpty(t, result.Warehouse)
	assert.Equal(t, result.Warehouse.Location.Address, input.Address)
	assert.Equal(t, result.Warehouse.CreatedBy, middleware.FakeUserID)
	assert.Equal(t, result.Warehouse.LastUpdatedBy, middleware.FakeUserID)
}

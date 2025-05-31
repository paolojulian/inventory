package stock_uc_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"paolojulian.dev/inventory/domain/product"
	"paolojulian.dev/inventory/domain/stock"
	"paolojulian.dev/inventory/domain/user"
	"paolojulian.dev/inventory/domain/warehouse"
	"paolojulian.dev/inventory/pkg/id"
	"paolojulian.dev/inventory/tests/factory"
	"paolojulian.dev/inventory/usecase/stock_uc"
)

type mockStockRepo struct {
	stockEntries map[string]*stock.StockEntry
}

func (repo *mockStockRepo) GetByID(ctx context.Context, stockEntryID string) (*stock.StockEntry, error) {
	existingStockEntry, exists := repo.stockEntries[stockEntryID]

	if !exists {
		return nil, stock_uc.ErrStockEntryNotFound
	}

	return existingStockEntry, nil
}

type mockProductRepo struct {
	products map[string]*product.ProductSummary
}

func (repo *mockProductRepo) GetSummary(ctx context.Context, productID string) (*product.ProductSummary, error) {
	existingProduct, exists := repo.products[productID]

	if !exists {
		return nil, stock_uc.ErrProductIDNotFound
	}

	return existingProduct, nil
}

type mockWarehouseRepo struct {
	warehouses map[string]*warehouse.WarehouseSummary
}

func (repo *mockWarehouseRepo) GetSummary(ctx context.Context, warehouseID string) (*warehouse.WarehouseSummary, error) {
	existingWarehouse, exists := repo.warehouses[warehouseID]

	if !exists {
		return nil, stock_uc.ErrWarehouseIDNotFound
	}

	return existingWarehouse, nil
}

type mockUserRepo struct {
	users map[string]*user.UserSummary
}

func (repo *mockUserRepo) GetSummary(ctx context.Context, userID string) (*user.UserSummary, error) {
	existingUser, exists := repo.users[userID]

	if !exists {
		return nil, stock_uc.ErrUserNotFound
	}

	return existingUser, nil
}

var mockStockEntryID = id.NewUUID()

func generateMockData() (stockRepo *mockStockRepo, productRepo *mockProductRepo, warehouseRepo *mockWarehouseRepo, userRepo *mockUserRepo) {
	fakeStockEntry := factory.NewTestStockEntry()
	fakeStockEntry.ID = mockStockEntryID
	stockRepo = &mockStockRepo{
		stockEntries: map[string]*stock.StockEntry{
			fakeStockEntry.ID: fakeStockEntry,
		},
	}

	fakeProduct := factory.NewTestProductSummary()
	fakeProduct.ID = fakeStockEntry.ProductID
	productRepo = &mockProductRepo{
		products: map[string]*product.ProductSummary{
			fakeProduct.ID: fakeProduct,
		},
	}

	fakeWarehouseSummary := factory.NewTestWarehouseSummary()
	fakeWarehouseSummary.ID = fakeStockEntry.WarehouseID
	warehouseRepo = &mockWarehouseRepo{
		warehouses: map[string]*warehouse.WarehouseSummary{
			fakeWarehouseSummary.ID: fakeWarehouseSummary,
		},
	}

	fakeUserSummary := factory.NewTestUserSummary()
	fakeUserSummary.ID = fakeStockEntry.UserID
	userRepo = &mockUserRepo{
		users: map[string]*user.UserSummary{
			fakeUserSummary.ID: fakeUserSummary,
		},
	}

	return stockRepo, productRepo, warehouseRepo, userRepo
}

func TestGetStockEntry_Success(t *testing.T) {
	stockRepo, productRepo, warehouseRepo, userRepo := generateMockData()
	uc := stock_uc.NewGetStockEntryUseCase(stockRepo, productRepo, warehouseRepo, userRepo)

	result, err := uc.Execute(context.Background(), mockStockEntryID)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}

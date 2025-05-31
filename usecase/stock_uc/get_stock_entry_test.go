package stock_uc_test

import (
	"context"

	"paolojulian.dev/inventory/domain/product"
	"paolojulian.dev/inventory/domain/stock"
	"paolojulian.dev/inventory/domain/user"
	"paolojulian.dev/inventory/domain/warehouse"
)

type mockStockRepo struct {
	stockEntries map[string]*stock.StockEntry
}

func (repo *mockStockRepo) GetByID(ctx context.Context, stockEntryID string) (stock.StockEntry, error)

type mockProductRepo struct {
	products map[string]*product.Product
}

func (repo *mockProductRepo) GetSummary(ctx context.Context, productID string) (product.ProductSummary, error)

type mockWarehouseRepo struct {
	warehouses map[string]*warehouse.Warehouse
}

func (repo *mockWarehouseRepo) GetSummary(ctx context.Context, productID string) (warehouse.WarehouseSummary, error)

type mockUserRepo struct {
	users map[string]*user.User
}

func (repo *mockUserRepo) GetSummary(ctx context.Context, productID string) (user.UserSummary, error)

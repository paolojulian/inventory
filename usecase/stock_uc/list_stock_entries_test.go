package stock_uc_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"paolojulian.dev/inventory/config"
	"paolojulian.dev/inventory/domain/stock"
	"paolojulian.dev/inventory/tests/factory"
	"paolojulian.dev/inventory/usecase/stock_uc"
)

type MockListStockEntriesRepo struct {
	stockEntries []*stock.StockEntry
}

func (repo *MockListStockEntriesRepo) GetList(ctx context.Context, limit int) ([]*stock.StockEntry, error) {
	return repo.stockEntries, nil
}

func TestListStockEntries_Success(t *testing.T) {
	repo := &MockListStockEntriesRepo{
		stockEntries: []*stock.StockEntry{
			factory.NewTestStockEntry(),
		},
	}
	uc := stock_uc.NewListStockEntriesUseCase(repo)
	input := &stock_uc.ListStockEntriesInput{
		Limit: config.IntPointer(1),
	}

	result, err := uc.Execute(context.Background(), input)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}

package inventory_uc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"paolojulian.dev/inventory/domain/inventory"
	"paolojulian.dev/inventory/domain/product"
	"paolojulian.dev/inventory/domain/warehouse"
)

type MockGetAllCurrentStockData struct {
}

type MockGetAllCurrentStockRepository struct {
	stocks []*inventory.InventoryItem
}

func (m *MockGetAllCurrentStockRepository) GetAllCurrentStock(ctx context.Context, warehouseID string) ([]*inventory.InventoryItem, error) {
	return m.stocks, nil
}

func TestGetAllCurrentStock(t *testing.T) {
	ctx := context.Background()
	repo := &MockGetAllCurrentStockRepository{
		stocks: []*inventory.InventoryItem{
			inventory.NewInventoryItem(
				&product.Product{
					ID:   "1",
					SKU:  product.SKU("1234567890"),
					Name: "Product 1",
					Price: product.Money{
						Cents: 100,
					},
				},
				&warehouse.Warehouse{
					ID:   "1",
					Name: "Warehouse 1",
				},
				10,
			),
		},
	}

	uc := NewGetAllCurrentStockUseCase(repo)

	output, err := uc.Execute(ctx, GetAllCurrentStockInput{})

	assert.NoError(t, err)
	assert.Equal(t, 1, len(output.Stocks))
}

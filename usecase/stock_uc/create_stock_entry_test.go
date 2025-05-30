package stock_uc_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"paolojulian.dev/inventory/config"
	"paolojulian.dev/inventory/domain/stock"
	"paolojulian.dev/inventory/pkg/id"
	"paolojulian.dev/inventory/usecase/stock_uc"
)

type MockCreateStockEntryWarehouseRepo struct {
	stockEntries         map[string]*stock.StockEntry
	existingProductIDs   []string
	existingWarehouseIDs []string
}

func (uc *MockCreateStockEntryWarehouseRepo) DoesProductIDExist(productID string) bool {
	return config.Includes(uc.existingProductIDs, productID)
}

func (uc *MockCreateStockEntryWarehouseRepo) DoesWarehouseIDExist(warehouseID string) bool {
	return config.Includes(uc.existingWarehouseIDs, warehouseID)
}

func (uc *MockCreateStockEntryWarehouseRepo) CreateStockEntry(ctx context.Context, stockEntry *stock.StockEntry) (*stock.StockEntry, error) {
	if !uc.DoesProductIDExist(stockEntry.ProductID) {
		return nil, stock_uc.ErrProductIDNotFound
	}

	if !uc.DoesWarehouseIDExist(stockEntry.WarehouseID) {
		return nil, stock_uc.ErrWarehouseIDNotFound
	}

	return stockEntry, nil
}

func TestCreateStockEntry__Success(t *testing.T) {
	userID := "test-user-id"
	initialProductIDs := []string{
		id.NewUUID(),
	}
	initialWarehouseIDs := []string{
		id.NewUUID(),
	}
	mockRepo := &MockCreateStockEntryWarehouseRepo{
		existingProductIDs:   initialProductIDs,
		existingWarehouseIDs: initialWarehouseIDs,
		stockEntries:         map[string]*stock.StockEntry{},
	}
	uc := stock_uc.NewCreateStockEntryUseCase(mockRepo)

	input := &stock_uc.StockEntryInput{
		QuantityDelta: 2,
		Reason:        string(stock.ReasonSale),
		ProductID:     initialProductIDs[0],
		WarehouseID:   initialWarehouseIDs[0],
	}

	_, err := uc.Execute(context.Background(), input, userID)
	assert.NoError(t, err)
}

func TestCreateStockEntry__NonExistingProductID(t *testing.T) {
	userID := "test-user-id"
	initialProductIDs := []string{
		id.NewUUID(),
	}
	initialWarehouseIDs := []string{
		id.NewUUID(),
	}
	mockRepo := &MockCreateStockEntryWarehouseRepo{
		existingProductIDs:   initialProductIDs,
		existingWarehouseIDs: initialWarehouseIDs,
		stockEntries:         map[string]*stock.StockEntry{},
	}
	uc := stock_uc.NewCreateStockEntryUseCase(mockRepo)

	input := &stock_uc.StockEntryInput{
		QuantityDelta: 2,
		Reason:        string(stock.ReasonSale),
		ProductID:     id.NewUUID(),
		WarehouseID:   initialWarehouseIDs[0],
	}

	_, err := uc.Execute(context.Background(), input, userID)
	assert.Error(t, err)
	assert.Equal(t, err, stock_uc.ErrProductIDNotFound)
}

func TestCreateStockEntry__NonExistingWarehouseID(t *testing.T) {
	userID := "test-user-id"
	initialProductIDs := []string{
		id.NewUUID(),
	}
	initialWarehouseIDs := []string{
		id.NewUUID(),
	}
	mockRepo := &MockCreateStockEntryWarehouseRepo{
		existingProductIDs:   initialProductIDs,
		existingWarehouseIDs: initialWarehouseIDs,
		stockEntries:         map[string]*stock.StockEntry{},
	}
	uc := stock_uc.NewCreateStockEntryUseCase(mockRepo)

	input := &stock_uc.StockEntryInput{
		QuantityDelta: 2,
		Reason:        string(stock.ReasonSale),
		ProductID:     initialProductIDs[0],
		WarehouseID:   id.NewUUID(),
	}

	_, err := uc.Execute(context.Background(), input, userID)
	assert.Error(t, err)
	assert.Equal(t, err, stock_uc.ErrWarehouseIDNotFound)
}

func TestCreateStockEntry__InvalidReason(t *testing.T) {
	userID := "test-user-id"
	initialProductIDs := []string{
		id.NewUUID(),
	}
	initialWarehouseIDs := []string{
		id.NewUUID(),
	}
	mockRepo := &MockCreateStockEntryWarehouseRepo{
		existingProductIDs:   initialProductIDs,
		existingWarehouseIDs: initialWarehouseIDs,
		stockEntries:         map[string]*stock.StockEntry{},
	}
	uc := stock_uc.NewCreateStockEntryUseCase(mockRepo)

	input := &stock_uc.StockEntryInput{
		QuantityDelta: 2,
		Reason:        "random-reason",
		ProductID:     initialProductIDs[0],
		WarehouseID:   id.NewUUID(),
	}

	_, err := uc.Execute(context.Background(), input, userID)
	assert.Error(t, err)
	assert.Equal(t, err, stock_uc.ErrInvalidReason)
}

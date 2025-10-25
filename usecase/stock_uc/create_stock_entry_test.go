package stock_uc

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"paolojulian.dev/inventory/config"
	"paolojulian.dev/inventory/domain/stock"
)

type MockCreateStockEntryRepo struct {
}

func (repo *MockCreateStockEntryRepo) CreateStockEntry(ctx context.Context, stockEntry *stock.StockEntry) (*stock.StockEntry, error) {
	return stockEntry, nil
}

func TestCreateStockEntry_ValidInput(t *testing.T) {
	repo := &MockCreateStockEntryRepo{}
	uc := NewCreateStockEntryUseCase(repo)

	input := StockEntryInput{
		QuantityDelta:      10,
		Reason:             string(stock.ReasonSale),
		ProductID:          "123",
		WarehouseID:        "456",
		SupplierPriceCents: config.IntPointer(1000),
		StorePriceCents:    config.IntPointer(1500),
		ExpiryDate:         config.TimePointer(time.Now().Add(time.Hour * 24 * 30)),
		ReorderDate:        config.TimePointer(time.Now().Add(time.Hour * 24 * 30)),
	}

	result, err := uc.Execute(context.Background(), &input, "123")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.StockEntry)
	assert.Equal(t, input.QuantityDelta, result.StockEntry.QuantityDelta)
	assert.Equal(t, input.ProductID, result.StockEntry.ProductID)
	assert.Equal(t, input.WarehouseID, result.StockEntry.WarehouseID)
	assert.Equal(t, input.SupplierPriceCents, result.StockEntry.SupplierPriceCents)
	assert.Equal(t, input.StorePriceCents, result.StockEntry.StorePriceCents)
	assert.Equal(t, input.ExpiryDate, result.StockEntry.ExpiryDate)
	assert.Equal(t, input.ReorderDate, result.StockEntry.ReorderDate)
	assert.Equal(t, stock.ReasonSale, result.StockEntry.Reason)
}

func TestCreateStockEntry_InvalidReason(t *testing.T) {
	repo := &MockCreateStockEntryRepo{}
	uc := NewCreateStockEntryUseCase(repo)

	input := StockEntryInput{
		QuantityDelta:      10,
		Reason:             "invalid",
		ProductID:          "123",
		WarehouseID:        "456",
		SupplierPriceCents: config.IntPointer(1000),
		StorePriceCents:    config.IntPointer(1500),
		ExpiryDate:         config.TimePointer(time.Now().Add(time.Hour * 24 * 30)),
		ReorderDate:        config.TimePointer(time.Now().Add(time.Hour * 24 * 30)),
	}

	result, err := uc.Execute(context.Background(), &input, "123")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, ErrInvalidReason, err)
}

func TestCreateStockEntry_InvalidSupplierPrice(t *testing.T) {
	repo := &MockCreateStockEntryRepo{}
	uc := NewCreateStockEntryUseCase(repo)

	input := StockEntryInput{
		QuantityDelta:      10,
		Reason:             string(stock.ReasonSale),
		ProductID:          "123",
		WarehouseID:        "456",
		SupplierPriceCents: config.IntPointer(-1),
	}

	result, err := uc.Execute(context.Background(), &input, "123")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, ErrInvalidSupplierPrice, err)
}

func TestCreateStockEntry_InvalidStorePrice(t *testing.T) {
	repo := &MockCreateStockEntryRepo{}
	uc := NewCreateStockEntryUseCase(repo)

	input := StockEntryInput{
		QuantityDelta:      10,
		Reason:             string(stock.ReasonSale),
		ProductID:          "123",
		WarehouseID:        "456",
		SupplierPriceCents: config.IntPointer(1000),
		StorePriceCents:    config.IntPointer(-1),
	}

	result, err := uc.Execute(context.Background(), &input, "123")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, ErrInvalidStorePrice, err)
}

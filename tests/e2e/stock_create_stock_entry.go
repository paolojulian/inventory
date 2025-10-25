package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"paolojulian.dev/inventory/domain/stock"
	"paolojulian.dev/inventory/infrastructure/postgres"
	"paolojulian.dev/inventory/interfaces/rest"
	"paolojulian.dev/inventory/tests/factory"
)

type StockEntryResponse struct {
	Message    string            `json:"message"`
	StockEntry *stock.StockEntry `json:"stock_entry"`
}

func TestStock_CreateStockEntry(t *testing.T) {
	// ================================
	// == Setup ==
	// ================================
	os.Setenv("APP_ENV", "test")
	gin.SetMode(gin.TestMode)
	bootstrap := rest.Bootstrap()
	ctx := context.Background()

	if err := cleanupTables(ctx, bootstrap.DB); err != nil {
		t.Fatalf("Failed to cleanup tables: %v", err)
	}
	defer cleanupTables(ctx, bootstrap.DB)
	defer bootstrap.DBCleanup()

	// After cleanup, re-run migrations and populate initial data so default records exist
	if err := postgres.MigrateSchema(bootstrap.DB); err != nil {
		t.Fatalf("failed to re-run migrations: %v", err)
	}
	if err := postgres.PopulateInitialData(bootstrap.DB); err != nil {
		t.Fatalf("failed to populate initial data: %v", err)
	}

	productRepo := postgres.NewProductRepository(bootstrap.DB)

	p1 := factory.NewTestProduct()
	p1.SKU = "INVSKU-001"
	createdP1, err := productRepo.Save(ctx, p1)
	if err != nil {
		t.Fatalf("failed to create product 1: %v", err)
	}

	p2 := factory.NewTestProduct()
	p2.SKU = "INVSKU-002"
	_, err = productRepo.Save(ctx, p2)
	if err != nil {
		t.Fatalf("failed to create product 2: %v", err)
	}

	// ================================
	// == Test ==
	// ================================
	input := map[string]interface{}{
		"product_id":           createdP1.ID,
		"warehouse_id":         "550e8400-e29b-41d4-a716-446655440000",
		"quantity_delta":       10,
		"reason":               "sale",
		"supplier_price_cents": 1000,
		"store_price_cents":    1500,
	}
	body, _ := json.Marshal(input)
	req := httptest.NewRequest("POST", "/stock/create-stock-entry", bytes.NewReader(body))
	w := httptest.NewRecorder()
	bootstrap.Router.ServeHTTP(w, req)

	// Parse response
	var resp StockEntryResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "Stock entry created successfully", resp.Message)
	assert.NotNil(t, resp.StockEntry)
	assert.Equal(t, input["product_id"], resp.StockEntry.ProductID)
	assert.Equal(t, input["warehouse_id"], resp.StockEntry.WarehouseID)
	assert.Equal(t, input["quantity_delta"], resp.StockEntry.QuantityDelta)
	assert.Equal(t, input["reason"], resp.StockEntry.Reason)
	assert.Equal(t, input["supplier_price_cents"], resp.StockEntry.SupplierPriceCents)
	assert.Equal(t, input["store_price_cents"], resp.StockEntry.StorePriceCents)
}

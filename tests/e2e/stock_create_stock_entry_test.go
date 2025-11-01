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

	userRepo := postgres.NewUserRepository(bootstrap.DB)
	adminUser, err := userRepo.FindByUsername(ctx, "admin")
	if err != nil {
		t.Fatalf("failed to find admin user: %v", err)
	}
	if adminUser == nil {
		t.Fatalf("admin user not found")
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

	req := httptest.NewRequest("POST", "/stock-entries", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	bootstrap.Router.ServeHTTP(w, req)

	// Parse response
	var resp StockEntryResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotNil(t, resp.StockEntry)
	assert.Equal(t, input["product_id"], resp.StockEntry.ProductID)
	assert.Equal(t, input["warehouse_id"], resp.StockEntry.WarehouseID)
	assert.Equal(t, input["quantity_delta"], resp.StockEntry.QuantityDelta)
	assert.Equal(t, stock.StockReason(input["reason"].(string)), resp.StockEntry.Reason)
	if resp.StockEntry.SupplierPriceCents != nil {
		assert.Equal(t, input["supplier_price_cents"], *resp.StockEntry.SupplierPriceCents)
	}
	if resp.StockEntry.StorePriceCents != nil {
		assert.Equal(t, input["store_price_cents"], *resp.StockEntry.StorePriceCents)
	}
}

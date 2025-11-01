package e2e

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"paolojulian.dev/inventory/infrastructure/postgres"
	"paolojulian.dev/inventory/interfaces/rest"
	"paolojulian.dev/inventory/tests/factory"
)

type InventoryListResponse struct {
	Message string      `json:"message"`
	Stocks  interface{} `json:"stocks"`
	Pager   struct {
		TotalItems  int `json:"total_items"`
		TotalPages  int `json:"total_pages"`
		CurrentPage int `json:"current_page"`
		PageSize    int `json:"page_size"`
	} `json:"pager"`
}

func TestInventory_GetAllStock_WithPager(t *testing.T) {
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

	// Seed two products and stock entries to ensure non-zero counts
	productRepo := postgres.NewProductRepository(bootstrap.DB)
	stockRepo := postgres.NewStockRepository(bootstrap.DB)

	p1 := factory.NewTestProduct()
	p1.SKU = "INVSKU-001"
	createdP1, err := productRepo.Save(ctx, p1)
	if err != nil {
		t.Fatalf("failed to create product 1: %v", err)
	}

	p2 := factory.NewTestProduct()
	p2.SKU = "INVSKU-002"
	createdP2, err := productRepo.Save(ctx, p2)
	if err != nil {
		t.Fatalf("failed to create product 2: %v", err)
	}

	se1 := factory.NewTestStockEntry()
	se1.ProductID = createdP1.ID
	se1.WarehouseID = "550e8400-e29b-41d4-a716-446655440000"
	// Use default admin user created by PopulateInitialData
	adminUser, err := postgres.NewUserRepository(bootstrap.DB).FindByUsername(ctx, "admin")
	if err != nil || adminUser == nil {
		t.Fatalf("failed to resolve default admin user: %v", err)
	}
	se1.UserID = adminUser.ID
	if _, err := stockRepo.CreateStockEntry(ctx, se1); err != nil {
		t.Fatalf("failed to create stock entry 1: %v", err)
	}

	se2 := factory.NewTestStockEntry()
	se2.ProductID = createdP2.ID
	se2.WarehouseID = "550e8400-e29b-41d4-a716-446655440000"
	se2.UserID = adminUser.ID
	if _, err := stockRepo.CreateStockEntry(ctx, se2); err != nil {
		t.Fatalf("failed to create stock entry 2: %v", err)
	}

	// Exercise route with pager
	req := httptest.NewRequest("GET", "/inventory/all-stock?page=1&size=20", nil)
	w := httptest.NewRecorder()
	bootstrap.Router.ServeHTTP(w, req)

	// Parse response
	var resp InventoryListResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	// Assertions
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "All stock retrieved successfully", resp.Message)
	assert.Equal(t, 1, resp.Pager.CurrentPage)
	assert.Equal(t, 20, resp.Pager.PageSize)
	// There should be at least 2 items after seeding
	assert.GreaterOrEqual(t, resp.Pager.TotalItems, 2)
	// Total pages should be consistent with total items and page size
	expectedTotalPages := (resp.Pager.TotalItems + resp.Pager.PageSize - 1) / resp.Pager.PageSize
	assert.Equal(t, expectedTotalPages, resp.Pager.TotalPages)
}

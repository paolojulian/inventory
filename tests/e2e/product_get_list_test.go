package e2e

import (
	"context"
	"encoding/json"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"paolojulian.dev/inventory/domain/product"
	"paolojulian.dev/inventory/infrastructure/postgres"
	"paolojulian.dev/inventory/interfaces/rest"
	paginationShared "paolojulian.dev/inventory/shared/pagination"
	"paolojulian.dev/inventory/tests/factory"
)

type ProductListResponse struct {
	Message  string                        `json:"message"`
	Products *[]product.Product            `json:"products"` // or your actual Product struct
	Pager    *paginationShared.PagerOutput `json:"pager"`    // or your actual PagerOutput struct
}

func TestGetList_ValidInput(t *testing.T) {
	// Set test environment
	os.Setenv("APP_ENV", "test")

	gin.SetMode(gin.TestMode)
	bootstrap := rest.Bootstrap()
	ctx := context.Background()

	// Clean up existing data first
	if err := cleanupTables(ctx, bootstrap.DB); err != nil {
		t.Fatalf("Failed to cleanup tables: %v", err)
	}
	defer cleanupTables(ctx, bootstrap.DB)
	defer bootstrap.DBCleanup()

	productRepo := postgres.NewProductRepository(bootstrap.DB)

	// Initialize test data
	product := factory.NewTestProduct()
	product.SKU = "TESTSKU456112"
	if _, err := productRepo.Save(ctx, product); err != nil {
		log.Fatal("Failed to create a mock data 1", err)
	}
	product2 := factory.NewTestProduct()
	product2.SKU = "TESTSKU456212"
	if _, err := productRepo.Save(ctx, product2); err != nil {
		log.Fatal("Failed to create a mock data 2", err)
	}

	// Prepare query parameters like frontend
	req := httptest.NewRequest("GET", "/products?page=1&size=10", nil)
	w := httptest.NewRecorder()

	bootstrap.Router.ServeHTTP(w, req)

	var response ProductListResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")

	// Assert
	assert.Equal(t, "Products fetched successfully", response.Message)
	assert.Len(t, *response.Products, 2)
	assert.Equal(t, 2, response.Pager.TotalItems)
	assert.Equal(t, 1, response.Pager.CurrentPage)
	assert.Equal(t, 10, response.Pager.PageSize)
	assert.Equal(t, 1, response.Pager.TotalPages)
}

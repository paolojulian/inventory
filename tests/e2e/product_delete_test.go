package e2e

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	productDomain "paolojulian.dev/inventory/domain/product"
	"paolojulian.dev/inventory/infrastructure/postgres"
	"paolojulian.dev/inventory/interfaces/rest"
	"paolojulian.dev/inventory/pkg/id"
)

func TestDeleteProduct_ValidInput(t *testing.T) {
	gin.SetMode(gin.TestMode)
	bootstrap := rest.Bootstrap()
	ctx := context.Background()

	// Use the repository directly to insert test data
	productRepo := postgres.NewProductRepository(bootstrap.DB)
	product := &productDomain.Product{
		ID:          id.NewUUID(), // or id.NewUUID()
		SKU:         "TESTSKU123",
		Name:        "Sample Product",
		Description: "This is a test product.",
		Price:       productDomain.Money{Cents: 4999},
		IsActive:    true,
	}
	created, err := productRepo.Save(ctx, product)
	assert.NoError(t, err)
	assert.NotNil(t, created)

	// Now that the product exists, we can test the DELETE request
	req := httptest.NewRequest("DELETE", "/products/"+created.ID, nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	bootstrap.Router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNoContent, w.Code)

	// Cleanup
	cleanupTables(context.Background(), bootstrap.DB)
	bootstrap.DBCleanup()
}

func TestDeleteProduct_InvalidProductID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	bootstrap := rest.Bootstrap()
	ctx := context.Background()

	// Use the repository directly to insert test data
	productRepo := postgres.NewProductRepository(bootstrap.DB)
	product := &productDomain.Product{
		ID:          id.NewUUID(), // or id.NewUUID()
		SKU:         "TESTSKU123",
		Name:        "Sample Product",
		Description: "This is a test product.",
		Price:       productDomain.Money{Cents: 4999},
		IsActive:    true,
	}
	created, err := productRepo.Save(ctx, product)
	assert.NoError(t, err)
	assert.NotNil(t, created)

	// Now that the product exists, we can test the DELETE request
	req := httptest.NewRequest("DELETE", "/products", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	bootstrap.Router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)

	// Cleanup
	cleanupTables(context.Background(), bootstrap.DB)
	bootstrap.DBCleanup()
}

func TestDeleteProduct_NonExistingID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	bootstrap := rest.Bootstrap()

	newId := id.NewUUID()
	// Delete non-existing
	req := httptest.NewRequest("DELETE", "/products/"+newId, nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	bootstrap.Router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNoContent, w.Code)

	// Cleanup
	cleanupTables(context.Background(), bootstrap.DB)
	bootstrap.DBCleanup()
}

package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	productDomain "paolojulian.dev/inventory/domain/product"
	"paolojulian.dev/inventory/infrastructure/postgres"
	"paolojulian.dev/inventory/interfaces/rest"
	"paolojulian.dev/inventory/tests/factory"
)

func TestUpdateProduct_ValidInput(t *testing.T) {
	gin.SetMode(gin.TestMode)
	bootstrap := rest.Bootstrap()
	ctx := context.Background()

	// == Create test data==
	productRepo := postgres.NewProductRepository(bootstrap.DB)
	product := factory.NewTestProduct()

	created, err := productRepo.Save(ctx, product)
	assert.NoError(t, err)
	assert.NotNil(t, created)
	// == End test data

	// == Start test ==
	input := map[string]interface{}{
		"name":        "New Name",
		"description": "New description",
		"price":       5999,
	}

	body, _ := json.Marshal(input)

	req := httptest.NewRequest("PUT", "/products/"+created.ID, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	bootstrap.Router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	updatedProduct, err := productRepo.GetByID(ctx, product.ID)
	assert.NoError(t, err)

	assert.Equal(t, "New Name", updatedProduct.Name)
	assert.Equal(t, productDomain.Description("New description"), updatedProduct.Description)
	assert.Equal(t, 5999, updatedProduct.Price.Cents)

	// == Cleanup ==
	cleanupTables(context.Background(), bootstrap.DB)
	bootstrap.DBCleanup()
}

func TestUpdateProduct_PartialUpdateShouldRemoveOtherFields(t *testing.T) {
	gin.SetMode(gin.TestMode)
	bootstrap := rest.Bootstrap()
	ctx := context.Background()

	// == Start test data ==
	productRepo := postgres.NewProductRepository(bootstrap.DB)
	product := factory.NewTestProduct()

	if created, err := productRepo.Save(ctx, product); err != nil {
		assert.NoError(t, err)
		assert.NotNil(t, created)
	}
	// == End test data ==

	// == Start test ==
	input := map[string]interface{}{
		"price": 5999,
	}

	body, _ := json.Marshal(input)

	req := httptest.NewRequest("PUT", "/products/"+product.ID, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	bootstrap.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	updatedProduct, err := productRepo.GetByID(ctx, product.ID)
	assert.NoError(t, err)

	assert.Equal(t, "", updatedProduct.Name)
	assert.Equal(t, productDomain.Description(""), updatedProduct.Description)
	assert.Equal(t, 5999, updatedProduct.Price.Cents)

	// == Cleanup ==
	cleanupTables(context.Background(), bootstrap.DB)
	bootstrap.DBCleanup()
}

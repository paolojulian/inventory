package e2e

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"paolojulian.dev/inventory/infrastructure/postgres"
	"paolojulian.dev/inventory/interfaces/rest"
	"paolojulian.dev/inventory/tests/factory"
)

func TestActiveProduct_ValidInput(t *testing.T) {
	gin.SetMode(gin.TestMode)
	bootstrap := rest.Bootstrap()
	ctx := context.Background()

	// == Cleanup ==
	defer cleanupTables(context.Background(), bootstrap.DB)
	defer bootstrap.DBCleanup()

	// == Create test data==
	productRepo := postgres.NewProductRepository(bootstrap.DB)
	product := factory.NewTestProduct()
	product.IsActive = false

	created, err := productRepo.Save(ctx, product)
	assert.NoError(t, err)
	assert.NotNil(t, created)
	// == End test data

	// == Start test ==
	req := httptest.NewRequest("POST", "/products/"+created.ID+"/activate", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	bootstrap.Router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	updatedProduct, err := productRepo.GetByID(ctx, product.ID)
	assert.NoError(t, err)

	assert.Equal(t, true, updatedProduct.IsActive)
}

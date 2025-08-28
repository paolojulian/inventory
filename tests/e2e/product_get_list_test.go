package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"paolojulian.dev/inventory/infrastructure/postgres"
	"paolojulian.dev/inventory/interfaces/rest"
	"paolojulian.dev/inventory/tests/factory"
)

func TestGetList_ValidInput(t *testing.T) {
	gin.SetMode(gin.TestMode)
	bootstrap := rest.Bootstrap()
	ctx := context.Background()

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

	// Prepare input
	input := map[string]interface{}{
		"pager": map[string]interface{}{
			"page": 1,
			"size": 10,
		},
	}

	body, _ := json.Marshal(input)

	req := httptest.NewRequest("GET", "/products", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	bootstrap.Router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code, "Expected status code 200 OK")
	assert.Equal(t, 2, len(w.Body.String()), "Expected 2 products in response")
	assert.Equal(t, "Products fetched successfully", w.Body.String(), "Expected success message in response")
}

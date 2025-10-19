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
	"paolojulian.dev/inventory/interfaces/rest"
)

func TestCreateProduct__ValidInput(t *testing.T) {
	// Set test environment
	os.Setenv("APP_ENV", "test")
	gin.SetMode(gin.TestMode)
	bootstrap := rest.Bootstrap()
	ctx := context.Background()

	// == Cleanup ==
	if err := cleanupTables(ctx, bootstrap.DB); err != nil {
		t.Fatalf("Failed to cleanup tables: %v", err)
	}
	defer bootstrap.DBCleanup()

	// Prepare input
	input := map[string]interface{}{
		"name":        "Sample Product",
		"description": "This is a test product.",
		"sku":         "TESTSKU123",
		"price":       4999,
	}

	body, _ := json.Marshal(input)

	// Perform request
	req := httptest.NewRequest("POST", "/products", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	bootstrap.Router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.NoError(t, err)
	assert.NotEmpty(t, response["product"])

}

func TestCreateProduct__NoInput(t *testing.T) {
	gin.SetMode(gin.TestMode)
	bootstrap := rest.Bootstrap()

	// Prepare input
	input := map[string]interface{}{}

	body, _ := json.Marshal(input)

	// Perform request
	req := httptest.NewRequest("POST", "/products", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	bootstrap.Router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.NoError(t, err)

	// Cleanup
	cleanupTables(context.Background(), bootstrap.DB)
	bootstrap.DBCleanup()
}

func TestCreateProduct__IncompleteInput(t *testing.T) {
	gin.SetMode(gin.TestMode)
	bootstrap := rest.Bootstrap()

	// Prepare input
	input := map[string]interface{}{
		"name":        "Sample Product",
		"description": "This is a test product.",
		// "sku":         "TESTSKU123", sku is missing
		"price": 4999,
	}

	body, _ := json.Marshal(input)

	// Perform request
	req := httptest.NewRequest("POST", "/products", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	bootstrap.Router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.NoError(t, err)

	// Cleanup
	cleanupTables(context.Background(), bootstrap.DB)
	bootstrap.DBCleanup()
}

func TestCreateProduct__DuplicateSKU(t *testing.T) {
	gin.SetMode(gin.TestMode)
	bootstrap := rest.Bootstrap()

	// Prepare input
	input := map[string]interface{}{
		"name":        "Sample Product",
		"description": "This is a test product.",
		"sku":         "TESTSKU123",
		"price":       4999,
	}

	body, _ := json.Marshal(input)

	// Perform request
	req := httptest.NewRequest("POST", "/products", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	bootstrap.Router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)

	// Run the api again for the same SKU
	body2, _ := json.Marshal(input)

	// Perform request
	req2 := httptest.NewRequest("POST", "/products", bytes.NewReader(body2))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()

	bootstrap.Router.ServeHTTP(w2, req2)

	assert.Equal(t, http.StatusConflict, w2.Code)

	// Cleanup
	cleanupTables(context.Background(), bootstrap.DB)
	bootstrap.DBCleanup()
}

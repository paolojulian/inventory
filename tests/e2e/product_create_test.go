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
	"paolojulian.dev/inventory/interfaces/rest"
)

func TestCreateProduct__ValidInput(t *testing.T) {
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

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.NoError(t, err)
	assert.NotEmpty(t, response["product"])

	// Cleanup
	cleanupTables(context.Background(), bootstrap.DB)
	bootstrap.DBCleanup()
}

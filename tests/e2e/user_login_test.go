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
	"paolojulian.dev/inventory/infrastructure/postgres"
	"paolojulian.dev/inventory/interfaces/rest"
	"paolojulian.dev/inventory/tests/factory"
)

func TestLogin__Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	bootstrap := rest.Bootstrap()
	ctx := context.Background()
	defer bootstrap.DB.Close()

	// == Create test data==
	username := "test-user"
	password := "qwe123!"
	userRepo := postgres.NewUserRepository(bootstrap.DB)
	user := factory.NewTestUser(password)
	user.Username = username

	created, err := userRepo.Save(ctx, user)
	assert.NoError(t, err, "unexpected error while creating test data")
	assert.NotNil(t, created, "test user was not created")
	// == End test data

	// == Start test ==
	input := map[string]interface{}{
		"username": username,
		"password": password,
	}

	body, _ := json.Marshal(input)

	req := httptest.NewRequest("POST", "/auth/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	bootstrap.Router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response body
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "failed to parse response body")

	// Check that the token is returned in the response
	token, ok := response["token"].(string)
	assert.True(t, ok, "token field not found in response")
	assert.NotEmpty(t, token, "token should not be empty")
	assert.Greater(t, len(token), 20, "token should be a valid JWT")
}

func TestLogin__Fail(t *testing.T) {
	gin.SetMode(gin.TestMode)
	bootstrap := rest.Bootstrap()
	ctx := context.Background()
	defer bootstrap.DB.Close()

	// == Create test data==
	username := "test-user"
	password := "qwe123!"
	userRepo := postgres.NewUserRepository(bootstrap.DB)
	user := factory.NewTestUser(password)
	user.Username = username

	created, err := userRepo.Save(ctx, user)
	assert.NoError(t, err, "unexpected error while creating test data")
	assert.NotNil(t, created, "test user was not created")
	// == End test data

	// == Start test ==
	input := map[string]interface{}{
		"username": username,
		"password": "another-password",
	}

	body, _ := json.Marshal(input)

	req := httptest.NewRequest("POST", "/auth/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	bootstrap.Router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Empty(t, w.Body)
}

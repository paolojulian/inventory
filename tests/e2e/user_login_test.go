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
	cleanupTables(ctx, bootstrap.DB)
	defer bootstrap.DB.Close()

	// == Create test data==
	username := "test-user"
	password := "qwe123"
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

	// == Cleanup ==
	cleanupTables(ctx, bootstrap.DB)
	bootstrap.DBCleanup()
}

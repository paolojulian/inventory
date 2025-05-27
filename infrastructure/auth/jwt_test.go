package auth_test

import (
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"paolojulian.dev/inventory/infrastructure/auth"
)

func init() {
	// Set a test secret key
	os.Setenv("JWT_SECRET_KEY", "testsecretkey")
}

func TestNewAccessTokenAndParseToken(t *testing.T) {
	userID := "test-user-id"

	token, err := auth.NewAccessToken(userID)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedUserID, err := auth.ParseToken(string(token))
	assert.NoError(t, err)
	assert.Equal(t, userID, parsedUserID)
}

func TestParseToken_InvalidSignature(t *testing.T) {
	// token signed with a different secret
	claims := jwt.MapClaims{
		"sub": "someone-else",
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign with a wrong secret
	invalidToken, err := token.SignedString([]byte("wrong-secret"))
	assert.NoError(t, err)

	_, err = auth.ParseToken(invalidToken)
	assert.Error(t, err)
}

func TestParseToken_Malformed(t *testing.T) {
	_, err := auth.ParseToken("not-a-jwt")
	assert.Error(t, err)
}

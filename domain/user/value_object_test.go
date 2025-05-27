package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	userDomain "paolojulian.dev/inventory/domain/user"
)

func TestHashPassword_ReturnsValidHash(t *testing.T) {
	password := "qwe123!"

	hashed, err := userDomain.HashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hashed)
	assert.NotEqual(t, password, string(hashed)) // Hash must not equal raw password
}

func TestComparePassword_CorrectPassword(t *testing.T) {
	password := "secretpass"
	hashed, err := userDomain.HashPassword(password)

	assert.NoError(t, err)

	err = userDomain.ComparePassword(string(hashed), password)

	assert.NoError(t, err) // Should succeed with correct password
}

func TestComparePassword_WrongPassword(t *testing.T) {
	password := "correctpass"
	wrongPassword := "wrongpass"
	hashed, err := userDomain.HashPassword(password)

	assert.NoError(t, err)

	err = userDomain.ComparePassword(string(hashed), wrongPassword)

	assert.Error(t, err) // Should fail with incorrect password
}

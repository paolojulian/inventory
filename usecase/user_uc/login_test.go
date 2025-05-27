package user_test

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	userDomain "paolojulian.dev/inventory/domain/user"
	"paolojulian.dev/inventory/tests/factory"
	userUC "paolojulian.dev/inventory/usecase/user_uc"
)

// == Mocks ==

type MockUserRepo struct {
	users map[string]*userDomain.User
}

func (r *MockUserRepo) Login(ctx context.Context, input *userUC.LoginInput) error {
	existingUser, exists := r.users[input.Username]
	if !exists {
		return userUC.ErrUserNotFound
	}

	if err := userDomain.ComparePassword(existingUser.Password, input.Password); err != nil {
		return userUC.ErrWrongPassword
	}

	return nil
}

// == Tests ==
func TestLogin_Success(t *testing.T) {
	var password string = "qwe123!"
	testUser := factory.NewTestUser()
	hashedPassword, err := userDomain.HashPassword(password)
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
	testUser.Password = string(hashedPassword)

	repo := &MockUserRepo{
		users: map[string]*userDomain.User{
			testUser.UserName: testUser,
		},
	}
	uc := userUC.NewLoginUseCase(repo)

	input := &userUC.LoginInput{
		Username: testUser.UserName,
		Password: password,
	}

	result, err := uc.Execute(context.Background(), input)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.NotEmpty(t, result.Token)
}

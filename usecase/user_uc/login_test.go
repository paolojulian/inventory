package user_uc_test

import (
	"context"
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

func (r *MockUserRepo) FindByUsername(ctx context.Context, username string) (*userDomain.User, error) {
	existingUser, exists := r.users[username]
	if !exists {
		return nil, userUC.ErrUserNotFound
	}

	return existingUser, nil
}

// == Tests ==
func TestLogin_Success(t *testing.T) {
	var password string = "qwe123!"
	testUser := factory.NewTestUser(password)

	repo := &MockUserRepo{
		users: map[string]*userDomain.User{
			testUser.Username: testUser,
		},
	}
	uc := userUC.NewLoginUseCase(repo)

	input := &userUC.LoginInput{
		Username: testUser.Username,
		Password: password,
	}

	result, err := uc.Execute(context.Background(), input)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.NotEmpty(t, result.Token)
}

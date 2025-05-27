package user

import (
	"context"

	userDomain "paolojulian.dev/inventory/domain/user"
)

type LoginInput struct {
	Username string
	Password string
}

type LoginOutput struct {
	Token userDomain.UserAccessToken
}

type UserRepository interface {
	Login(ctx context.Context, input *LoginInput) error
}

type LoginUseCase struct {
	repo UserRepository
}

func NewLoginUseCase(repo UserRepository) *LoginUseCase {
	return &LoginUseCase{repo}
}

func (uc *LoginUseCase) Execute(ctx context.Context, input *LoginInput) (*LoginOutput, error) {
	err := uc.repo.Login(ctx, input)
	if err != nil {
		return nil, err
	}

	userAccessToken := userDomain.NewUserAccessToken()

	return &LoginOutput{
		Token: userAccessToken,
	}, nil
}

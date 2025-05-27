package user

import (
	"context"

	"paolojulian.dev/inventory/domain/user"
	"paolojulian.dev/inventory/infrastructure/auth"
)

type LoginInput struct {
	Username string
	Password string
}

type LoginOutput struct {
	Token auth.AccessToken
}

type UserRepository interface {
	Login(ctx context.Context, input *LoginInput) (*user.User, error)
}

type LoginUseCase struct {
	repo UserRepository
}

func NewLoginUseCase(repo UserRepository) *LoginUseCase {
	return &LoginUseCase{repo}
}

func (uc *LoginUseCase) Execute(ctx context.Context, input *LoginInput) (*LoginOutput, error) {
	user, err := uc.repo.Login(ctx, input)
	if err != nil {
		return nil, err
	}

	userAccessToken, err := auth.NewAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &LoginOutput{
		Token: userAccessToken,
	}, nil
}

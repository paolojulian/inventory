package user_uc

import (
	"context"
	"log"

	userDomain "paolojulian.dev/inventory/domain/user"
	"paolojulian.dev/inventory/infrastructure/auth"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginOutput struct {
	Token auth.AccessToken
}

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*userDomain.User, error)
}

type LoginUseCase struct {
	repo UserRepository
}

func NewLoginUseCase(repo UserRepository) *LoginUseCase {
	return &LoginUseCase{repo}
}

func (uc *LoginUseCase) Execute(ctx context.Context, input *LoginInput) (*LoginOutput, error) {
	user, err := uc.repo.FindByUsername(ctx, input.Username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	if err := userDomain.ComparePassword(user.Password, input.Password); err != nil {
		log.Printf("here")
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

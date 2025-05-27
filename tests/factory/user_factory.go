package factory

import (
	"paolojulian.dev/inventory/domain/user"
	"paolojulian.dev/inventory/pkg/id"
)

func NewTestUser() *user.User {
	return &user.User{
		ID:        id.NewUUID(),
		UserName:  "johndoe",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@email.com",
		Mobile:    "09279488654",
		Password:  "qweqweqwe123",
		Role:      user.AdminRole,
		IsActive:  true,
	}
}

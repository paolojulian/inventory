package factory

import (
	"log"

	"paolojulian.dev/inventory/domain/user"
	"paolojulian.dev/inventory/pkg/id"
)

func NewTestUser(password string) *user.User {
	hashedPassword, err := user.HashPassword(password)
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}

	firstName := "John"
	lastName := "Doe"
	email := "johndoe@email.com"
	mobile := "09279488654"

	return &user.User{
		ID:        id.NewUUID(),
		Username:  "johndoe",
		Password:  string(hashedPassword),
		FirstName: &firstName,
		LastName:  &lastName,
		Email:     &email,
		Mobile:    &mobile,
		Role:      string(user.AdminRole),
		IsActive:  true,
	}
}

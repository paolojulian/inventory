package user

import "paolojulian.dev/inventory/pkg/id"

type User struct {
	ID        string
	Username  string
	Password  string
	Role      string
	IsActive  bool
	Email     *string
	FirstName *string
	LastName  *string
	Mobile    *string
}

func NewUser(
	username string,
	password string,
	role UserRole,
	isActive bool,
	email *string,
	firstName *string,
	lastName *string,
	mobile *string,
) (*User, error) {
	hashed, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:        id.NewUUID(),
		Username:  username,
		Password:  string(hashed),
		Role:      string(role),
		IsActive:  isActive,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Mobile:    mobile,
	}, nil
}

type UserSummary struct {
	ID        string
	Username  string
	Role      string
	Email     *string
	FirstName *string
	LastName  *string
	Mobile    *string
}

package user

import "paolojulian.dev/inventory/pkg/id"

type User struct {
	ID        string  `json:"id"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	Role      string  `json:"role"`
	IsActive  bool    `json:"is_active"`
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Mobile    *string `json:"mobile,omitempty"`
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

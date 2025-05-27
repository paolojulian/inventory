package user

import "errors"

var (
	ErrWrongPassword   = errors.New("wrong password")
	ErrUserNotFound    = errors.New("user not found")
	ErrPasswordInvalid = errors.New("password is invalid")
)

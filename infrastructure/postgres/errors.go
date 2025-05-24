package postgres

import "errors"

var (
	ErrDatabaseUrlNotSet = errors.New("DATABASE_URL not set")
)
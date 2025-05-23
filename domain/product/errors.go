package product

import "errors"

var (
	ErrSKUMustBeAtLeast4Chars = errors.New("SKU must be at least 4 characters")
	ErrDescriptionTooLong     = errors.New("description is too long")
)

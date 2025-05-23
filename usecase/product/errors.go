package product

import "errors"

var (
	ErrSKUAlreadyExists      = errors.New("SKU already exists")
	ErrInvalidSKU            = errors.New("invalid SKU")
	ErrInvalidName           = errors.New("invalid name")
	ErrInvalidPrice          = errors.New("invalid price")
	ErrUnableToDeleteProduct = errors.New("unable to delete product")
	ErrDBError               = errors.New("unable to delete product")
)

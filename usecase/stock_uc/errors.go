package stock_uc

import "errors"

var (
	ErrInvalidReason       = errors.New("invalid reason")
	ErrProductIDNotFound   = errors.New("product id not found")
	ErrWarehouseIDNotFound = errors.New("warehouse id not found")
)

package stock_uc

import "errors"

var (
	ErrInvalidReason       = errors.New("invalid reason")
	ErrProductIDNotFound   = errors.New("product id not found")
	ErrWarehouseIDNotFound = errors.New("warehouse id not found")
	ErrUserNotFound        = errors.New("user not found")
	ErrStockEntryNotFound  = errors.New("stock entry not found")
)

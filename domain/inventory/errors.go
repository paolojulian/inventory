package inventory

import "errors"

var (
	ErrProductNotFound     = errors.New("product not found")
	ErrWarehouseNotFound   = errors.New("warehouse not found")
	ErrInvalidStockLevel   = errors.New("invalid stock level")
	ErrInsufficientStock   = errors.New("insufficient stock")
	ErrStockCalculationFailed = errors.New("stock calculation failed")
)

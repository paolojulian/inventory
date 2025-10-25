package stock

import "errors"

var (
	ErrQuantityDeltaZero    = errors.New("quantity delta cannot be zero")
	ErrReasonEmpty          = errors.New("reason cannot be empty")
	ErrInvalidSupplierPrice = errors.New("invalid supplier price")
	ErrInvalidStorePrice    = errors.New("invalid store price")
)

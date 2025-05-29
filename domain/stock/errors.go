package stock

import "errors"

var (
	ErrQuantityDeltaZero = errors.New("quantity delta cannot be zero")
	ErrReasonEmpty       = errors.New("reason cannot be empty")
)

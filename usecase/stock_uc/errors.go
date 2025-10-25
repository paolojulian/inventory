package stock_uc

import "errors"

var (
	ErrInvalidReason        = errors.New("invalid reason")
	ErrInvalidSupplierPrice = errors.New("invalid supplier price")
	ErrInvalidStorePrice    = errors.New("invalid store price")
	ErrInvalidExpiryDate    = errors.New("invalid expiry date")
	ErrProductIDNotFound    = errors.New("product id not found")
	ErrWarehouseIDNotFound  = errors.New("warehouse id not found")
	ErrUserNotFound         = errors.New("user not found")
	ErrStockEntryNotFound   = errors.New("stock entry not found")
)

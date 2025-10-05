package inventory

// StockStatus represents the status of stock levels
type StockStatus string

const (
	StatusInStock     StockStatus = "in_stock"
	StatusLowStock    StockStatus = "low_stock"
	StatusOutOfStock  StockStatus = "out_of_stock"
)

// IsValid checks if the stock status is valid
func (s StockStatus) IsValid() bool {
	switch s {
	case StatusInStock, StatusLowStock, StatusOutOfStock:
		return true
	default:
		return false
	}
}

// String returns the string representation of the stock status
func (s StockStatus) String() string {
	return string(s)
}

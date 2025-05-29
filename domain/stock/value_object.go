package stock

type StockReason string

const (
	ReasonRestock     StockReason = "restock"
	ReasonSale        StockReason = "sale"
	ReasonDamage      StockReason = "damage"
	ReasonTransferIn  StockReason = "transfer_in"
	ReasonTransferOut StockReason = "transfer_out"
	ReasonAdjustment  StockReason = "adjustment"
)

func IsValidStockReason(reason string) bool {
	switch StockReason(reason) {
	case
		ReasonRestock,
		ReasonSale,
		ReasonDamage,
		ReasonTransferIn,
		ReasonTransferOut,
		ReasonAdjustment:
		return true
	default:
		return false
	}
}

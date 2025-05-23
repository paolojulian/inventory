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

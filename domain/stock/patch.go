package stock

type StockEntryPatch struct {
	QuantityDelta *int
	Reason        *StockReason
	ProductID     *string
	WarehouseID   *string
	UserID        string
}

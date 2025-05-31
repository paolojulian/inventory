package warehouse

type Warehouse struct {
	ID            string
	Location      WarehouseLocation
	IsActive      bool
	CreatedBy     string
	LastUpdatedBy string
}

type WarehouseSummary struct {
	ID       string
	Location WarehouseLocation
}

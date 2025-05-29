package warehouse

type Warehouse struct {
	ID            string
	Location      WarehouseLocation
	IsActive      bool
	CreatedBy     string
	LastUpdatedBy string
}

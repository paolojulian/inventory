package warehouse

type WarehouseLocation struct {
	Address     string
	City        *string
	Region      *string
	Country     *string
	Coordinates *Coordinates
}

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

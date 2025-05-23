package product

type Product struct {
	ID       string
	SKU      SKU
	Name     string
	Price    Money
	IsActive bool
}

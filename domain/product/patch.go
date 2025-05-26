package product

type ProductPatch struct {
	ID          *string
	SKU         *SKU
	Name        *string
	Description *Description
	Price       *Money
	IsActive    *bool
}

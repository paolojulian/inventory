package product

import (
	"context"

	"paolojulian.dev/inventory/domain/product"
)

type CreateProductInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	SKU         string `json:"sku" binding:"required"`
	Price       int    `json:"price" binding:"required"` // Cents, e.g. PHP 49.99 = 4999
}

type CreateProductOutput struct {
	Product *product.Product
}

type ProductRepository interface {
	Save(ctx context.Context, product *product.Product) (*product.Product, error)
	ExistsBySKU(ctx context.Context, sku string) (bool, error)
}

type CreateProductUseCase struct {
	repo ProductRepository
}

func NewCreateProductUseCase(repo ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{repo}
}

func (uc *CreateProductUseCase) Execute(ctx context.Context, input CreateProductInput) (*CreateProductOutput, error) {
	// Validate input
	sku, err := product.NewSKU(input.SKU)
	if err != nil {
		return &CreateProductOutput{}, err
	}

	description, err := product.NewDescription(input.Description)
	if err != nil {
		return &CreateProductOutput{}, err
	}

	// Validate SKU
	exists, err := uc.repo.ExistsBySKU(ctx, string(sku))
	if err != nil {
		return &CreateProductOutput{}, err
	}
	if exists {
		return &CreateProductOutput{}, ErrSKUAlreadyExists
	}

	createdProduct := product.NewProduct(
		sku,
		input.Name,
		description,
		input.Price,
	)

	product, err := uc.repo.Save(ctx, createdProduct)
	if err != nil {
		return &CreateProductOutput{}, err
	}

	return &CreateProductOutput{product}, nil
}

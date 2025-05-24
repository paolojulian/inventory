package product

import (
	"context"

	"paolojulian.dev/inventory/domain/product"
	"paolojulian.dev/inventory/pkg/id"
)

type CreateProductInput struct {
	Name        string
	Description string
	SKU         string
	Price       int // Cents, e.g. PHP 49.99 = 4999
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

	createdProduct := &product.Product{
		ID:          id.NewULID(),
		SKU:         sku,
		Name:        input.Name,
		Description: description,
		Price:       product.Money{Cents: input.Price},
		IsActive:    true,
	}

	product, err := uc.repo.Save(ctx, createdProduct)
	if err != nil {
		return &CreateProductOutput{}, err
	}

	return &CreateProductOutput{product}, nil
}

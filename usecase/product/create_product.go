package product

import (
	"context"

	"paolojulian.dev/inventory/domain/product"
	"paolojulian.dev/inventory/pkg/id"
)

type CreateProductInput struct {
	Name  string
	SKU   string
	Price int // Cents, e.g. PHP 49.99 = 4999
}

type CreateProductOutput struct {
	ProductID string
}

type ProductRepository interface {
	Save(ctx context.Context, product *product.Product) error
	ExistsBySKU(ctx context.Context, sku string) (bool, error)
}

type CreateProductUseCase struct {
	repo ProductRepository
}

func NewCreateProductUseCase(repo ProductRepository) *CreateProductUseCase {
	return &CreateProductUseCase{repo}
}

func (uc *CreateProductUseCase) Execute(ctx context.Context, input CreateProductInput) (*CreateProductOutput, error) {
	sku := product.SKU(input.SKU)

	// Validate SKU
	exists, err := uc.repo.ExistsBySKU(ctx, string(sku))
	if err != nil {
		return &CreateProductOutput{}, err
	}
	if exists {
		return &CreateProductOutput{}, ErrSKUAlreadyExists
	}

	createdProduct := &product.Product{
		ID:       id.NewULID(),
		SKU:      sku,
		Name:     input.Name,
		Price:    product.Money{Cents: input.Price},
		IsActive: true,
	}

	if err := uc.repo.Save(ctx, createdProduct); err != nil {
		return &CreateProductOutput{}, err
	}

	return &CreateProductOutput{ProductID: createdProduct.ID}, nil
}

package product

import (
	"context"

	productDomain "paolojulian.dev/inventory/domain/product"
)

type UpdateProductBasicInput struct {
	Name        string
	Description string
	Price       int
}

type UpdateProductBasicOutput struct {
	Product *productDomain.Product
}

type UpdateProductBasicRepo interface {
	UpdateByID(ctx context.Context, productID string, product *productDomain.Product) (*productDomain.Product, error)
}

type UpdateProductBasicUseCase struct {
	repo UpdateProductBasicRepo
}

func NewUpdateProductBasicUseCase(repo UpdateProductBasicRepo) *UpdateProductBasicUseCase {
	return &UpdateProductBasicUseCase{repo}
}

func (uc *UpdateProductBasicUseCase) Execute(ctx context.Context, productID string, input UpdateProductBasicInput) (*UpdateProductBasicOutput, error) {
	// Parse and validate value objects
	description, err := productDomain.NewDescription(input.Description)
	if err != nil {
		return nil, err
	}

	updated := &productDomain.Product{
		Name:        input.Name,
		Description: description,
		Price:       productDomain.Money{Cents: input.Price},
		// Other fields are unchanged
	}

	newProduct, err := uc.repo.UpdateByID(ctx, productID, updated)
	if err != nil {
		return nil, err
	}

	return &UpdateProductBasicOutput{Product: newProduct}, nil
}

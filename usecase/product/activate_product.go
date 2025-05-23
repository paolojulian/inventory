package product

import (
	"context"

	productDomain "paolojulian.dev/inventory/domain/product"
)

type ActivateProductInput struct {
	ProductID string
}

type ActivateProductOutput struct {
	Product *productDomain.Product
}

type ActivateProductRepo interface {
	ActivateProductByID(ctx context.Context, productID string) (*productDomain.Product, error)
}

type ActivateProductUseCase struct {
	repo ActivateProductRepo
}

func NewActivateProductUseCase(repo ActivateProductRepo) *ActivateProductUseCase {
	return &ActivateProductUseCase{repo}
}

func (uc *ActivateProductUseCase) Execute(ctx context.Context, productID string) (*ActivateProductOutput, error) {
	newProduct, err := uc.repo.ActivateProductByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	return &ActivateProductOutput{Product: newProduct}, nil
}

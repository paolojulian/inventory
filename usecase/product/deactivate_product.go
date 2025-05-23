package product

import (
	"context"

	productDomain "paolojulian.dev/inventory/domain/product"
)

type DeactivateProductInput struct {
	ProductID string
}

type DeactivateProductOutput struct {
	Product *productDomain.Product
}

type DeactivateProductRepo interface {
	DeactivateByID(ctx context.Context, productID string) (*productDomain.Product, error)
}

type DeactivateProductUseCase struct {
	repo DeactivateProductRepo
}

func NewDeactivateProductUseCase(repo DeactivateProductRepo) *DeactivateProductUseCase {
	return &DeactivateProductUseCase{repo}
}

func (uc *DeactivateProductUseCase) Execute(ctx context.Context, productID string) (*DeactivateProductOutput, error) {
	newProduct, err := uc.repo.DeactivateByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	return &DeactivateProductOutput{Product: newProduct}, nil
}

package product_uc

import (
	"context"

	productDomain "paolojulian.dev/inventory/domain/product"
)

type UpdateProductBasicInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Price       *int    `json:"price"`
	SKU         *string `json:"sku"`
}

type UpdateProductBasicOutput struct {
	Product *productDomain.Product
}

type UpdateProductBasicRepo interface {
	UpdateByID(ctx context.Context, productID string, product *productDomain.ProductPatch) (*productDomain.Product, error)
}

type UpdateProductBasicUseCase struct {
	repo UpdateProductBasicRepo
}

func NewUpdateProductBasicUseCase(repo UpdateProductBasicRepo) *UpdateProductBasicUseCase {
	return &UpdateProductBasicUseCase{repo}
}

func (uc *UpdateProductBasicUseCase) Execute(ctx context.Context, productID string, input UpdateProductBasicInput) (*UpdateProductBasicOutput, error) {
	var description *productDomain.Description
	if input.Description != nil {
		newDescription, err := productDomain.NewDescription(*input.Description)
		if err != nil {
			return nil, err
		}
		description = &newDescription
	}

	var price *productDomain.Money
	if input.Price != nil {
		price = &productDomain.Money{
			Cents: *input.Price,
		}
	}

	var sku *productDomain.SKU
	if input.SKU != nil {
		newSku, err := productDomain.PtrSKUFromString(input.SKU)
		if err != nil {
			return nil, err
		}
		sku = newSku
	}

	updated := &productDomain.ProductPatch{
		Name:        input.Name,
		Description: description,
		Price:       price,
		SKU:         sku,
		// Other fields are unchanged
	}

	newProduct, err := uc.repo.UpdateByID(ctx, productID, updated)
	if err != nil {
		return nil, err
	}

	return &UpdateProductBasicOutput{Product: newProduct}, nil
}

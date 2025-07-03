package product_uc

import (
	"context"

	productDomain "paolojulian.dev/inventory/domain/product"
)

type GetProductListInput struct {
	Filter *productDomain.ProductFilter
	Sort   *productDomain.ProductSort
}

type GetProductListOutput struct {
	Products []productDomain.Product
}

type GetProductListRepo interface {
	GetList(ctx context.Context, filter *productDomain.ProductFilter, sort *productDomain.ProductSort) ([]productDomain.Product, error)
}

type GetProductListUseCase struct {
	repo GetProductListRepo
}

func NewGetProductListUseCase(repo GetProductListRepo) *GetProductListUseCase {
	return &GetProductListUseCase{repo}
}

func (uc *GetProductListUseCase) Execute(ctx context.Context, input GetProductListInput) (GetProductListOutput, error) {
	products, err := uc.repo.GetList(ctx, input.Filter, input.Sort)
	if err != nil {
		return GetProductListOutput{}, err
	}

	return GetProductListOutput{Products: products}, nil
}

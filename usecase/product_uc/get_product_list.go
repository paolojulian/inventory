package product_uc

import (
	"context"

	productDomain "paolojulian.dev/inventory/domain/product"
	paginationShared "paolojulian.dev/inventory/shared/pagination"
)

type GetProductListInput struct {
	Pager  paginationShared.PagerInput  `json:"pager"`
	Filter *productDomain.ProductFilter `json:"filter,omitempty"`
	Sort   *productDomain.ProductSort   `json:"sort,omitempty"`
}

type GetProductListOutput struct {
	Products *[]productDomain.Product
	Pager    *paginationShared.PagerOutput
}

type GetListOutput struct {
	Products []productDomain.Product
	Pager    paginationShared.PagerOutput
}

type GetProductListUseCase struct {
	repo GetProductListRepo
}

type GetProductListRepo interface {
	GetList(ctx context.Context, pager paginationShared.PagerInput, filter *productDomain.ProductFilter, sort *productDomain.ProductSort) (*productDomain.GetListOutput, error)
}

func NewGetProductListUseCase(repo GetProductListRepo) *GetProductListUseCase {
	return &GetProductListUseCase{repo}
}

func (uc *GetProductListUseCase) Execute(ctx context.Context, input GetProductListInput) (*GetProductListOutput, error) {
	result, err := uc.repo.GetList(ctx, input.Pager, input.Filter, input.Sort)
	if err != nil {
		return &GetProductListOutput{}, err
	}

	return &GetProductListOutput{Products: result.Products, Pager: result.Pager}, nil
}

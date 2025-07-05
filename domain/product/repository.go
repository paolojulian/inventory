package product

import (
	"context"

	paginationShared "paolojulian.dev/inventory/shared/pagination"
)

type GetListOutput struct {
	Products *[]Product
	Pager    *paginationShared.PagerOutput
}

type IProductRepository interface {
	GetList(ctx context.Context, pager *paginationShared.PagerInput, filter *ProductFilter, sort *ProductSort) (GetListOutput, error)
}

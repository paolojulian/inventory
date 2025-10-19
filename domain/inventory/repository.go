package inventory

import (
	"context"

	paginationShared "paolojulian.dev/inventory/shared/pagination"
)

type GetAllStockOutput struct {
	Stocks []*InventoryItem
	Pager  paginationShared.PagerOutput
}

type IInventoryRepository interface {
	GetAllStock(ctx context.Context, warehouseID string, pager paginationShared.PagerInput) (*GetAllStockOutput, error)
}

package stock_uc

import (
	"context"

	"paolojulian.dev/inventory/config"
	"paolojulian.dev/inventory/domain/stock"
)

type UpdateStockEntryInput struct {
	QuantityDelta *int    `json:"quantity_data,omit_empty"`
	Reason        *string `json:"reason,omit_empty"`
	ProductID     *string `json:"product_id,omit_empty"`
	WarehouseID   *string `json:"warehouse_id,omit_empty"`
	UserID        *string `json:"user_id,omit_empty"`
}

type UpdateStockEntryOutput struct {
	StockEntry *stock.StockEntry
}

type UpdateStockEntryRepo interface {
	UpdateByID(ctx context.Context, stockEntryID string, stockEntryPatch *stock.StockEntryPatch) (*stock.StockEntry, error)
}

type UpdateStockEntryUseCase struct {
	repo UpdateStockEntryRepo
}

func NewUpdateStockEntryUseCase(repo UpdateStockEntryRepo) *UpdateStockEntryUseCase {
	return &UpdateStockEntryUseCase{repo}
}

func (uc *UpdateStockEntryUseCase) Execute(ctx context.Context, stockEntryID string, input UpdateStockEntryInput, userID string) (*UpdateStockEntryOutput, error) {
	var reason stock.StockReason
	if input.Reason != nil {
		if !stock.IsValidStockReason(*input.Reason) {
			return nil, ErrInvalidReason
		}
		reason = stock.StockReason(*input.Reason)
	}

	updated := &stock.StockEntryPatch{
		QuantityDelta: config.IntPointer(*input.QuantityDelta),
		Reason:        &reason,
		ProductID:     config.StringPointer(*input.ProductID),
		WarehouseID:   config.StringPointer(*input.WarehouseID),
		UserID:        userID,
	}

	newStockEntry, err := uc.repo.UpdateByID(ctx, stockEntryID, updated)
	if err != nil {
		return nil, err
	}

	return &UpdateStockEntryOutput{StockEntry: newStockEntry}, nil
}

package stock_uc

import "context"

type DeleteStockEntryInput struct {
	StockEntryID string `json:"stock_entry_id" binding:"required"`
}

type DeleteStockEntryOutput struct {
	IsSuccess bool
}

type DeleteStockEntryRepository interface {
	Delete(ctx context.Context, stockEntryID string) error
}

type DeleteStockEntryUseCase struct {
	repo DeleteStockEntryRepository
}

func NewDeleteStockEntryUseCase(repo DeleteStockEntryRepository) *DeleteStockEntryUseCase {
	return &DeleteStockEntryUseCase{repo}
}

func (uc *DeleteStockEntryUseCase) Execute(ctx context.Context, input DeleteStockEntryInput) (*DeleteStockEntryOutput, error) {
	if err := uc.repo.Delete(ctx, input.StockEntryID); err != nil {
		return &DeleteStockEntryOutput{IsSuccess: false}, err
	}

	return &DeleteStockEntryOutput{IsSuccess: true}, nil
}

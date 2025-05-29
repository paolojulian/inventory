package warehouse_uc

import "context"

type DeleteWarehouseInput struct {
	WarehouseID string `json:"warehouse_id" binding:"required"`
}

type DeleteWarehouseOutput struct {
	IsSuccess bool
}

type DeleteWarehouseRepository interface {
	Delete(ctx context.Context, productID string) error
}

type DeleteWarehouseUseCase struct {
	repo DeleteWarehouseRepository
}

func NewDeleteWarehouseUseCase(repo DeleteWarehouseRepository) *DeleteWarehouseUseCase {
	return &DeleteWarehouseUseCase{repo}
}

func (uc *DeleteWarehouseUseCase) Execute(ctx context.Context, input DeleteWarehouseInput) (*DeleteWarehouseOutput, error) {
	if err := uc.repo.Delete(ctx, input.WarehouseID); err != nil {
		return &DeleteWarehouseOutput{IsSuccess: false}, err
	}

	return &DeleteWarehouseOutput{IsSuccess: true}, nil
}

package inventory_uc

import (
	"context"

	"paolojulian.dev/inventory/domain/inventory"
)

type GetInventorySummaryInput struct {
	// No input needed - uses default warehouse
}

type GetInventorySummaryOutput struct {
	Summary *inventory.InventorySummary `json:"summary"`
}

type GetInventorySummaryRepository interface {
	GetInventorySummary(ctx context.Context) (*inventory.InventorySummary, error)
}

type GetInventorySummaryUseCase struct {
	repo GetInventorySummaryRepository
}

func NewGetInventorySummaryUseCase(repo GetInventorySummaryRepository) *GetInventorySummaryUseCase {
	return &GetInventorySummaryUseCase{repo}
}

func (uc *GetInventorySummaryUseCase) Execute(ctx context.Context, input GetInventorySummaryInput) (*GetInventorySummaryOutput, error) {
	summary, err := uc.repo.GetInventorySummary(ctx)
	if err != nil {
		return nil, err
	}

	return &GetInventorySummaryOutput{Summary: summary}, nil
}

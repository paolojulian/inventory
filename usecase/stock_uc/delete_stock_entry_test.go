package stock_uc_test

import (
	"context"
	"errors"
	"testing"

	"paolojulian.dev/inventory/usecase/stock_uc"
)

// --- Mock Repository ---

type MockDeleteStockEntryRepo struct {
	isDeleted  bool
	shouldFail bool
}

func (r *MockDeleteStockEntryRepo) Delete(ctx context.Context, stockEntryID string) error {
	if r.shouldFail {
		r.isDeleted = false
		return errors.New("unable to delete stockEntry")
	}

	r.isDeleted = true
	return nil
}

// --- Tests ---

func TestDeleteStockEntry_ValidInput(t *testing.T) {
	repo := &MockDeleteStockEntryRepo{}
	uc := stock_uc.NewDeleteStockEntryUseCase(repo)

	input := stock_uc.DeleteStockEntryInput{
		StockEntryID: "some-valid-id",
	}

	result, err := uc.Execute(context.Background(), input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsSuccess {
		t.Fatalf("expected deletion to be successful")
	}
}

func TestDeleteStockEntry_RepoError(t *testing.T) {
	repo := &MockDeleteStockEntryRepo{
		shouldFail: true,
	}
	uc := stock_uc.NewDeleteStockEntryUseCase(repo)

	input := stock_uc.DeleteStockEntryInput{
		StockEntryID: "some-valid-id",
	}

	result, err := uc.Execute(context.Background(), input)
	if err == nil {
		t.Fatalf("expected deletion to fail")
	}
	if result.IsSuccess {
		t.Fatalf("expected deletion to fail")
	}
}

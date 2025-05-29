package warehouse_uc_test

import (
	"context"
	"errors"
	"testing"

	"paolojulian.dev/inventory/usecase/warehouse_uc"
)

// --- Mock Repository ---

type MockWarehouseRepo struct {
	isDeleted  bool
	shouldFail bool
}

func (r *MockWarehouseRepo) Delete(ctx context.Context, warehouseID string) error {
	if r.shouldFail {
		r.isDeleted = false
		return errors.New("unable to delete product")
	}

	r.isDeleted = true
	return nil
}

// --- Tests ---

func TestDeleteProduct_ValidInput(t *testing.T) {
	repo := &MockWarehouseRepo{}
	uc := warehouse_uc.NewDeleteWarehouseUseCase(repo)

	input := warehouse_uc.DeleteWarehouseInput{
		WarehouseID: "some-valid-id",
	}

	result, err := uc.Execute(context.Background(), input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsSuccess {
		t.Fatalf("expected deletion to be successful")
	}
}

func TestDeleteProduct_RepoError(t *testing.T) {
	repo := &MockWarehouseRepo{
		shouldFail: true,
	}
	uc := warehouse_uc.NewDeleteWarehouseUseCase(repo)

	input := warehouse_uc.DeleteWarehouseInput{
		WarehouseID: "some-valid-id",
	}

	result, err := uc.Execute(context.Background(), input)
	if err == nil {
		t.Fatalf("expected deletion to fail")
	}
	if result.IsSuccess {
		t.Fatalf("expected deletion to fail")
	}
}

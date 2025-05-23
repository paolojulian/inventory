package product_test

import (
	"context"
	"errors"
	"testing"

	productUC "paolojulian.dev/inventory/usecase/product"
)

// --- Mock Repository ---

type MockDeleteProductRepo struct {
	isDeleted  bool
	shouldFail bool
}

func (r *MockDeleteProductRepo) Delete(ctx context.Context, productID string) error {
	if r.shouldFail {
		r.isDeleted = false
		return errors.New("unable to delete product")
	}

	r.isDeleted = true
	return nil
}

// --- Tests ---

func TestDeleteProduct_ValidInput(t *testing.T) {
	repo := &MockDeleteProductRepo{}
	uc := productUC.NewDeleteProductUseCase(repo)

	input := productUC.DeleteProductInput{
		ProductID: "some-valid-id",
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
	repo := &MockDeleteProductRepo{
		shouldFail: true,
	}
	uc := productUC.NewDeleteProductUseCase(repo)

	input := productUC.DeleteProductInput{
		ProductID: "some-valid-id",
	}

	result, err := uc.Execute(context.Background(), input)
	if err == nil {
		t.Fatalf("expected deletion to fail")
	}
	if result.IsSuccess {
		t.Fatalf("expected deletion to fail")
	}
}

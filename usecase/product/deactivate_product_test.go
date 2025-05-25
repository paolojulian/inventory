package product_test

import (
	"context"
	"testing"

	productDomain "paolojulian.dev/inventory/domain/product"
	productUC "paolojulian.dev/inventory/usecase/product"
)

// --- Mock Repo ---

type MockDeactivateProductRepo struct {
	products map[string]*productDomain.Product
}

func (r *MockDeactivateProductRepo) DeactivateProductByID(ctx context.Context, productID string) (*productDomain.Product, error) {
	if existingProduct, exists := r.products[productID]; exists {
		existingProduct.IsActive = false

		return existingProduct, nil
	}

	return nil, productUC.ErrProductNotFound
}

// --- Tests ---

func TestDeactivateProduct__Success(t *testing.T) {
	repo := &MockDeactivateProductRepo{
		products: map[string]*productDomain.Product{
			"existing-product-id": {
				ID:          "existing-product-id",
				Name:        "In-active Product",
				Description: productDomain.Description(""),
				Price:       productDomain.Money{Cents: 1000},
				IsActive:    false,
			},
		},
	}

	uc := productUC.NewDeactivateProductUseCase(repo)

	result, err := uc.Execute(context.Background(), "existing-product-id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
		return
	}

	if result.Product.IsActive {
		t.Fatalf("Expected %s, Got %t", "false", !result.Product.IsActive)
		return
	}
}

func TestDeactivateProduct__Fail(t *testing.T) {
	repo := &MockDeactivateProductRepo{
		products: map[string]*productDomain.Product{
			"existing-product-id": {
				ID:          "existing-product-id",
				Name:        "In-active Product",
				Description: productDomain.Description(""),
				Price:       productDomain.Money{Cents: 1000},
				IsActive:    false,
			},
		},
	}

	uc := productUC.NewDeactivateProductUseCase(repo)

	_, err := uc.Execute(context.Background(), "non-existing-product-id")
	if err == nil {
		t.Fatal("expected error but nothing given")
	}
}

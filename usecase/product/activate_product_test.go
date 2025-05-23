package product_test

import (
	"context"
	"testing"

	productDomain "paolojulian.dev/inventory/domain/product"
	productUC "paolojulian.dev/inventory/usecase/product"
)

// --- Mock Repo ---

type MockActivateProductRepo struct {
	products map[string]*productDomain.Product
}

func (r *MockActivateProductRepo) ActivateProductByID(ctx context.Context, productID string) (*productDomain.Product, error) {
	if existingProduct, exists := r.products[productID]; exists {
		existingProduct.IsActive = true

		return existingProduct, nil
	}

	return nil, productUC.ErrProductNotFound
}

// --- Tests ---

func TestActivateProduct__Success(t *testing.T) {
	repo := &MockActivateProductRepo{
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

	uc := productUC.NewActivateProductUseCase(repo)

	result, err := uc.Execute(context.Background(), "existing-product-id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
		return
	}

	if !result.Product.IsActive {
		t.Fatalf("Expected %s, Got %t", "true", result.Product.IsActive)
		return
	}
}

func TestActivateProduct__Fail(t *testing.T) {
	repo := &MockActivateProductRepo{
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

	uc := productUC.NewActivateProductUseCase(repo)

	_, err := uc.Execute(context.Background(), "non-existing-product-id")
	if err == nil {
		t.Fatal("expected error but nothing given")
	}
}

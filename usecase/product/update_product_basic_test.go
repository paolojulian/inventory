package product_test

import (
	"context"
	"testing"

	productDomain "paolojulian.dev/inventory/domain/product"
	productUC "paolojulian.dev/inventory/usecase/product"
)

// --- Mock Repo ---

type MockUpdateProductBasicRepo struct {
	updated  *productDomain.Product
	products map[string]*productDomain.Product
}

func (r *MockActivateProductRepo) UpdateByID(ctx context.Context, productID string, product *productDomain.Product) (*productDomain.Product, error) {
	if existingProduct, exists := r.products[productID]; exists {
		existingProduct.Name = product.Name
		existingProduct.Description = product.Description
		existingProduct.Price = product.Price

		return existingProduct, nil
	}

	return nil, productUC.ErrProductNotFound
}

// --- Tests ---

func TestUpdateProductBasic__ValidInput(t *testing.T) {
	repo := &MockActivateProductRepo{
		products: map[string]*productDomain.Product{
			"existing-product-id": {
				ID:          "existing-product-id",
				Name:        "Old Name",
				Description: productDomain.Description("Old Description"),
				Price:       productDomain.Money{Cents: 1000},
			},
		},
	}

	uc := productUC.NewUpdateProductBasicUseCase(repo)

	input := productUC.UpdateProductBasicInput{
		Name:        "New Name",
		Description: "New Description",
		Price:       2000,
	}

	result, err := uc.Execute(context.Background(), "existing-product-id", input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
		return
	}

	if result.Product.Name != "New Name" {
		t.Fatalf("Expected %s, Got %s", "New Name", result.Product.Name)
		return
	}

	if result.Product.Description != "New Description" {
		t.Fatalf("Expected %s, Got %s", "New Description", result.Product.Description)
		return
	}
}

func TestUpdateProductBasic__ProductNotFound(t *testing.T) {
	repo := &MockActivateProductRepo{
		products: map[string]*productDomain.Product{
			"existing-product-id": {
				ID:          "existing-product-id",
				Name:        "Old Name",
				Description: productDomain.Description("Old Description"),
				Price:       productDomain.Money{Cents: 1000},
			},
		},
	}

	uc := productUC.NewUpdateProductBasicUseCase(repo)

	input := productUC.UpdateProductBasicInput{
		Name:        "New Name",
		Description: "New Description",
		Price:       2000,
	}

	_, err := uc.Execute(context.Background(), "non-existing-product-id", input)
	if err == nil {
		t.Fatal("expected error but nothing given")
	}
}

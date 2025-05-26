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

func (r *MockActivateProductRepo) UpdateByID(ctx context.Context, productID string, product *productDomain.ProductPatch) (*productDomain.Product, error) {
	if existingProduct, exists := r.products[productID]; exists {
		if product.Name != nil {
			existingProduct.Name = *product.Name
		}
		if product.Description != nil {
			existingProduct.Description = *product.Description
		}
		if product.Price != nil {
			existingProduct.Price = *product.Price
		}
		if product.SKU != nil {
			existingProduct.SKU = *product.SKU
		}

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
				SKU:         productDomain.SKU("TESTSKU-123"),
			},
		},
	}

	uc := productUC.NewUpdateProductBasicUseCase(repo)

	var name string = "New Name"
	var description string = "New Description"
	var price int = 2000

	input := productUC.UpdateProductBasicInput{
		Name:        &name,
		Description: &description,
		Price:       &price,
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

	var name string = "New Name"
	var description string = "New Description"
	var price int = 2000

	input := productUC.UpdateProductBasicInput{
		Name:        &name,
		Description: &description,
		Price:       &price,
	}

	_, err := uc.Execute(context.Background(), "non-existing-product-id", input)
	if err == nil {
		t.Fatal("expected error but nothing given")
	}
}

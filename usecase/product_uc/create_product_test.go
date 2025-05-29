package product_uc_test

import (
	"context"
	"strings"
	"testing"

	productDomain "paolojulian.dev/inventory/domain/product"
	"paolojulian.dev/inventory/pkg/id"
	productUC "paolojulian.dev/inventory/usecase/product_uc"
)

// --- Mocks ---

type MockCreateProductRepo struct {
	saved        *productDomain.Product
	existingSKUs map[string]bool
}

func (r *MockCreateProductRepo) Save(ctx context.Context, product *productDomain.Product) (*productDomain.Product, error) {
	savedProduct := &productDomain.Product{
		ID:          id.NewUUID(),
		SKU:         product.SKU,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		IsActive:    true,
	}
	r.saved = savedProduct

	return savedProduct, nil
}

func (r *MockCreateProductRepo) ExistsBySKU(ctx context.Context, sku string) (bool, error) {
	return r.existingSKUs[sku], nil
}

// --- Tests ---

func TestCreateProduct_ValidInput(t *testing.T) {
	repo := &MockCreateProductRepo{existingSKUs: make(map[string]bool)}
	uc := productUC.NewCreateProductUseCase(repo)

	input := productUC.CreateProductInput{
		Name:  "Test Product",
		SKU:   "TSHIRT-LG-RED",
		Price: 19999,
	}

	result, err := uc.Execute(context.Background(), input)
	if err != nil {
		t.Fatalf("unexpected error: %v ", err)
	}

	if repo.saved == nil || repo.saved.SKU != "TSHIRT-LG-RED" {
		t.Fatalf("expected product to be saved with SKU 'TSHIRT-LG-RED'")
	}

	if result.Product == nil {
		t.Fatalf("expected a generated product")
	}

	if result.Product.ID == "" {
		t.Fatalf("expected a generated product ID")
	}
}

func TestCreateProduct_DescriptionTooLong(t *testing.T) {
	repo := &MockCreateProductRepo{existingSKUs: make(map[string]bool)}
	uc := productUC.NewCreateProductUseCase(repo)

	longDescription := strings.Repeat("x", 3001)
	input := productUC.CreateProductInput{
		Name:        "Test Product",
		Description: longDescription,
		SKU:         "TSHIRT-LG-RED",
		Price:       19999,
	}

	_, err := uc.Execute(context.Background(), input)
	if err == nil {
		t.Fatalf("expected description too long, got nil")
	}
}

func TestCreateProduct_SKUExists(t *testing.T) {
	repo := &MockCreateProductRepo{existingSKUs: map[string]bool{"TSHIRT-LG-RED": true}}
	uc := productUC.NewCreateProductUseCase(repo)

	input := productUC.CreateProductInput{
		Name:  "Test Product",
		SKU:   "TSHIRT-LG-RED",
		Price: 19999,
	}

	_, err := uc.Execute(context.Background(), input)
	if err == nil {
		t.Fatalf("expected SKU already exists error, got nil")
	}
}

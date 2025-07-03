package product_uc_test

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	productDomain "paolojulian.dev/inventory/domain/product"
	productUC "paolojulian.dev/inventory/usecase/product_uc"
)

// --- Mock Repo ---

type MockGetProductListRepo struct {
	products []productDomain.Product
}

func (r *MockGetProductListRepo) GetList(ctx context.Context, filter *productDomain.ProductFilter, sort *productDomain.ProductSort) ([]productDomain.Product, error) {
	if filter != nil {
		if filter.SearchText != nil {
			filteredProducts := []productDomain.Product{}
			for _, product := range r.products {
				if strings.Contains(product.Name, *filter.SearchText) {
					filteredProducts = append(filteredProducts, product)
				}
			}
			r.products = filteredProducts
		}

		if filter.IsActive != nil {
			filteredProducts := []productDomain.Product{}
			for _, product := range r.products {
				if product.IsActive == *filter.IsActive {
					filteredProducts = append(filteredProducts, product)
				}
			}
			r.products = filteredProducts
		}
	}

	return r.products, nil
}

func TestGetProductList__ValidInput(t *testing.T) {
	repo := &MockGetProductListRepo{
		products: []productDomain.Product{
			*productDomain.NewProduct(productDomain.SKU("SKU-1"), "Product 1", productDomain.Description("Description 1"), 100),
			*productDomain.NewProduct(productDomain.SKU("SKU-2"), "Product 2", productDomain.Description("Description 2"), 200),
			*productDomain.NewProduct(productDomain.SKU("SKU-3"), "Product 3", productDomain.Description("Description 3"), 300),
		},
	}

	uc := productUC.NewGetProductListUseCase(repo)

	input := productUC.GetProductListInput{}

	output, err := uc.Execute(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(output.Products))
	assert.Equal(t, productDomain.SKU("SKU-1"), output.Products[0].SKU)
	assert.Equal(t, productDomain.SKU("SKU-2"), output.Products[1].SKU)
	assert.Equal(t, productDomain.SKU("SKU-3"), output.Products[2].SKU)
}

func TestGetProductList__Filter__ValidInput(t *testing.T) {
	repo := &MockGetProductListRepo{
		products: []productDomain.Product{
			*productDomain.NewProduct(productDomain.SKU("SKU-1"), "Product 1", productDomain.Description("Description 1"), 100),
			*productDomain.NewProduct(productDomain.SKU("SKU-2"), "Product 2", productDomain.Description("Description 2"), 200),
			*productDomain.NewProduct(productDomain.SKU("SKU-3"), "Product 3", productDomain.Description("Description 3"), 300),
		},
	}

	uc := productUC.NewGetProductListUseCase(repo)

	searchText := "Product 1"

	input := productUC.GetProductListInput{
		Filter: &productDomain.ProductFilter{
			SearchText: &searchText,
		},
	}

	output, err := uc.Execute(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(output.Products))
	assert.Equal(t, productDomain.SKU("SKU-1"), output.Products[0].SKU)
}

func TestGetProductList__Filter__Status__ValidInput(t *testing.T) {
	inactiveProduct := productDomain.NewProduct(productDomain.SKU("SKU-1"), "Product 1", productDomain.Description("Description 1"), 100)
	inactiveProduct.IsActive = false

	repo := &MockGetProductListRepo{
		products: []productDomain.Product{
			*productDomain.NewProduct(productDomain.SKU("SKU-1"), "Product 1", productDomain.Description("Description 1"), 100),
			*productDomain.NewProduct(productDomain.SKU("SKU-2"), "Product 2", productDomain.Description("Description 2"), 200),
			*productDomain.NewProduct(productDomain.SKU("SKU-3"), "Product 3", productDomain.Description("Description 3"), 300),
			*inactiveProduct,
		},
	}

	uc := productUC.NewGetProductListUseCase(repo)

	isActive := true
	input := productUC.GetProductListInput{
		Filter: &productDomain.ProductFilter{
			IsActive: &isActive,
		},
	}

	output, err := uc.Execute(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(output.Products))
	assert.Equal(t, productDomain.SKU("SKU-1"), output.Products[0].SKU)
	assert.Equal(t, productDomain.SKU("SKU-2"), output.Products[1].SKU)
	assert.Equal(t, productDomain.SKU("SKU-3"), output.Products[2].SKU)
}

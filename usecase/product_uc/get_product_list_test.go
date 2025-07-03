package product_uc_test

import (
	"context"
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	productDomain "paolojulian.dev/inventory/domain/product"
	paginationShared "paolojulian.dev/inventory/shared/pagination"
	productUC "paolojulian.dev/inventory/usecase/product_uc"
)

// --- Mock Repo ---

type MockGetProductListData struct {
	products []productDomain.Product
}
type MockGetListOutput struct {
	Products []productDomain.Product
	Pager    paginationShared.PagerOutput
}

func (r *MockGetProductListData) GetList(ctx context.Context, pager paginationShared.PagerInput, filter *productDomain.ProductFilter, sort *productDomain.ProductSort) (*productDomain.GetListOutput, error) {
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

	totalItems := len(r.products)
	pageSize := pager.PageSize
	if totalItems < pager.PageSize {
		pageSize = totalItems
	}

	if pager.PageSize < totalItems {
		r.products = r.products[pager.PageSize:]
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))

	return &productDomain.GetListOutput{
		Products: &r.products,
		Pager: &paginationShared.PagerOutput{
			TotalItems:  totalItems,
			TotalPages:  totalPages,
			CurrentPage: pager.PageNumber,
			PageSize:    pageSize,
		},
	}, nil
}

func TestGetProductList__ValidInput(t *testing.T) {
	repo := &MockGetProductListData{
		products: []productDomain.Product{
			*productDomain.NewProduct(productDomain.SKU("SKU-1"), "Product 1", productDomain.Description("Description 1"), 100),
			*productDomain.NewProduct(productDomain.SKU("SKU-2"), "Product 2", productDomain.Description("Description 2"), 200),
			*productDomain.NewProduct(productDomain.SKU("SKU-3"), "Product 3", productDomain.Description("Description 3"), 300),
		},
	}

	uc := productUC.NewGetProductListUseCase(repo)

	input := productUC.GetProductListInput{
		Pager: paginationShared.PagerInput{
			PageNumber: 1,
			PageSize:   10,
		},
	}

	output, err := uc.Execute(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(*output.Products))
	assert.Equal(t, productDomain.SKU("SKU-1"), (*output.Products)[0].SKU)
	assert.Equal(t, productDomain.SKU("SKU-2"), (*output.Products)[1].SKU)
	assert.Equal(t, productDomain.SKU("SKU-3"), (*output.Products)[2].SKU)

	assert.Equal(t, 1, output.Pager.TotalPages)
	assert.Equal(t, 1, output.Pager.CurrentPage)
	assert.Equal(t, 3, output.Pager.PageSize)
	assert.Equal(t, 3, output.Pager.TotalItems)
}

func TestGetProductList__Filter__ValidInput(t *testing.T) {
	repo := &MockGetProductListData{
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
	assert.Equal(t, 1, len(*output.Products))
	assert.Equal(t, productDomain.SKU("SKU-1"), (*output.Products)[0].SKU)
}

func TestGetProductList__Filter__Status__ValidInput(t *testing.T) {
	inactiveProduct := productDomain.NewProduct(productDomain.SKU("SKU-1"), "Product 1", productDomain.Description("Description 1"), 100)
	inactiveProduct.IsActive = false

	repo := &MockGetProductListData{
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
	assert.Equal(t, 3, len(*output.Products))
	assert.Equal(t, productDomain.SKU("SKU-1"), (*output.Products)[0].SKU)
	assert.Equal(t, productDomain.SKU("SKU-2"), (*output.Products)[1].SKU)
	assert.Equal(t, productDomain.SKU("SKU-3"), (*output.Products)[2].SKU)
}

func TestGetProductList__Pager__ValidInput(t *testing.T) {
	repo := &MockGetProductListData{
		products: []productDomain.Product{
			*productDomain.NewProduct(productDomain.SKU("SKU-1"), "Product 1", productDomain.Description("Description 1"), 100),
			*productDomain.NewProduct(productDomain.SKU("SKU-2"), "Product 2", productDomain.Description("Description 2"), 200),
			*productDomain.NewProduct(productDomain.SKU("SKU-3"), "Product 3", productDomain.Description("Description 3"), 300),
			*productDomain.NewProduct(productDomain.SKU("SKU-4"), "Product 4", productDomain.Description("Description 4"), 400),
			*productDomain.NewProduct(productDomain.SKU("SKU-5"), "Product 5", productDomain.Description("Description 5"), 500),
			*productDomain.NewProduct(productDomain.SKU("SKU-6"), "Product 6", productDomain.Description("Description 6"), 600),
		},
	}

	uc := productUC.NewGetProductListUseCase(repo)

	input := productUC.GetProductListInput{
		Pager: paginationShared.PagerInput{
			PageNumber: 1,
			PageSize:   3,
		},
	}

	output, err := uc.Execute(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(*output.Products))
	assert.Equal(t, productDomain.SKU("SKU-4"), (*output.Products)[0].SKU)
	assert.Equal(t, productDomain.SKU("SKU-5"), (*output.Products)[1].SKU)
	assert.Equal(t, productDomain.SKU("SKU-6"), (*output.Products)[2].SKU)

	assert.Equal(t, 2, output.Pager.TotalPages, "Total pages should be 2")
	assert.Equal(t, 1, output.Pager.CurrentPage, "Current page should be 1")
	assert.Equal(t, 3, output.Pager.PageSize, "Page size should be 3")
	assert.Equal(t, 6, output.Pager.TotalItems, "Total items should be 6")
}

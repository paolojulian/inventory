package product

import "context"

type DeleteProductInput struct {
	ProductID string
}

type DeleteProductOutput struct {
	IsSuccess bool
}

type DeleteProductRepository interface {
	Delete(ctx context.Context, productID string) error
}

type DeleteProductUseCase struct {
	repo DeleteProductRepository
}

func NewDeleteProductUseCase(repo DeleteProductRepository) *DeleteProductUseCase {
	return &DeleteProductUseCase{repo}
}

func (uc *DeleteProductUseCase) Execute(ctx context.Context, input DeleteProductInput) (*DeleteProductOutput, error) {
	if err := uc.repo.Delete(ctx, input.ProductID); err != nil {
		return &DeleteProductOutput{IsSuccess: false}, err
	}

	return &DeleteProductOutput{IsSuccess: true}, nil
}

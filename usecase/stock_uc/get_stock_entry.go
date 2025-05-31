package stock_uc

import (
	"context"

	"paolojulian.dev/inventory/domain/product"
	"paolojulian.dev/inventory/domain/stock"
	"paolojulian.dev/inventory/domain/user"
	"paolojulian.dev/inventory/domain/warehouse"
)

type GetStockEntryInput struct {
	StockEntryID string
}

type GetStockEntryOutput struct {
	StockEntry *stock.StockEntry
	Product    *product.ProductSummary
	Warehouse  *warehouse.WarehouseSummary
	User       *user.UserSummary
}

type stockRepository interface {
	GetByID(ctx context.Context, stockEntryID string) (stock.StockEntry, error)
}

type productRepository interface {
	GetSummary(ctx context.Context, productID string) (product.ProductSummary, error)
}

type warehouseRepository interface {
	GetSummary(ctx context.Context, warehouseID string) (warehouse.WarehouseSummary, error)
}

type userRepository interface {
	GetSummary(ctx context.Context, userID string) (user.UserSummary, error)
}

type GetStockEntryUseCase struct {
	repo          stockRepository
	productRepo   productRepository
	warehouseRepo warehouseRepository
	userRepo      userRepository
}

func NewGetStockEntryUseCase(repo stockRepository, productRepo productRepository, warehouseRepo warehouseRepository, userRepo userRepository) *GetStockEntryUseCase {
	return &GetStockEntryUseCase{
		repo,
		productRepo,
		warehouseRepo,
		userRepo,
	}
}

func (uc *GetStockEntryUseCase) Execute(ctx context.Context, stockEntryID string) (*GetStockEntryOutput, error) {
	stockEntry, err := uc.repo.GetByID(ctx, stockEntryID)
	if err != nil {
		return nil, err
	}

	productSummary, err := uc.productRepo.GetSummary(ctx, stockEntry.ProductID)
	if err != nil {
		return nil, err
	}

	warehouseSummary, err := uc.warehouseRepo.GetSummary(ctx, stockEntry.WarehouseID)
	if err != nil {
		return nil, err
	}

	userSummary, err := uc.userRepo.GetSummary(ctx, stockEntry.UserID)
	if err != nil {
		return nil, err
	}

	return &GetStockEntryOutput{
		StockEntry: &stockEntry,
		Product:    &productSummary,
		Warehouse:  &warehouseSummary,
		User:       &userSummary,
	}, nil
}

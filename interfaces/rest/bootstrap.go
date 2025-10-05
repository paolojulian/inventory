package rest

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"paolojulian.dev/inventory/config"
	"paolojulian.dev/inventory/infrastructure/postgres"
	inventoryUC "paolojulian.dev/inventory/usecase/inventory_uc"
	productUC "paolojulian.dev/inventory/usecase/product_uc"
	stockUC "paolojulian.dev/inventory/usecase/stock_uc"
	"paolojulian.dev/inventory/usecase/user_uc"
)

type ProductHandlers struct {
	Activate   *productUC.ActivateProductUseCase
	Create     *productUC.CreateProductUseCase
	Deactivate *productUC.DeactivateProductUseCase
	Delete     *productUC.DeleteProductUseCase
	Update     *productUC.UpdateProductBasicUseCase
	GetList    *productUC.GetProductListUseCase
}

type AuthHandlers struct {
	Login *user_uc.LoginUseCase
}

type StockHandlers struct {
	Create  *stockUC.CreateStockEntryUseCase
	GetList *stockUC.ListStockEntriesUseCase
	Get     *stockUC.GetStockEntryUseCase
	Update  *stockUC.UpdateStockEntryUseCase
	Delete  *stockUC.DeleteStockEntryUseCase
}

type InventoryHandlers struct {
	GetCurrentStock *inventoryUC.GetCurrentStockUseCase
	GetAllStock     *inventoryUC.GetAllCurrentStockUseCase
	GetSummary      *inventoryUC.GetInventorySummaryUseCase
	GetLowStock     *inventoryUC.GetLowStockUseCase
	GetOutOfStock   *inventoryUC.GetOutOfStockUseCase
}

type Handlers struct {
	Product   *ProductHandlers
	Auth      *AuthHandlers
	Stock     *StockHandlers
	Inventory *InventoryHandlers
}

type Application struct {
	Router    *gin.Engine
	Handlers  *Handlers
	DBCleanup func()
	DB        *pgxpool.Pool
}

func Bootstrap() *Application {
	// Load configs this early
	config.LoadConfig()

	// Load the DB
	db, err := postgres.NewPool()
	if err != nil {
		log.Fatalf("failed to connect to the database server. error: %v", err)
	}

	if err := postgres.MigrateSchema(db); err != nil {
		log.Fatalf("unable to migrate schema: %v", err)
	}
	println("schema migrated successfully.")

	if err := postgres.PopulateInitialData(db); err != nil {
		log.Fatalf("unable to populate initial data: %v", err)
	}

	// Wire the repo to use cases
	productRepo := postgres.NewProductRepository(db)
	userRepo := postgres.NewUserRepository(db)
	stockRepo := postgres.NewStockRepository(db)
	warehouseRepo := postgres.NewWarehouseRepository(db)
	inventoryRepo := postgres.NewInventoryRepository(db)

	handlers := &Handlers{
		Product: &ProductHandlers{
			Activate:   productUC.NewActivateProductUseCase(productRepo),
			Create:     productUC.NewCreateProductUseCase(productRepo),
			Deactivate: productUC.NewDeactivateProductUseCase(productRepo),
			Delete:     productUC.NewDeleteProductUseCase(productRepo),
			Update:     productUC.NewUpdateProductBasicUseCase(productRepo),
			GetList:    productUC.NewGetProductListUseCase(productRepo),
		},
		Auth: &AuthHandlers{
			Login: user_uc.NewLoginUseCase(userRepo),
		},
		Stock: &StockHandlers{
			Create:  stockUC.NewCreateStockEntryUseCase(stockRepo),
			GetList: stockUC.NewListStockEntriesUseCase(stockRepo),
			Get:     stockUC.NewGetStockEntryUseCase(stockRepo, productRepo, warehouseRepo, userRepo),
			Update:  stockUC.NewUpdateStockEntryUseCase(stockRepo),
			Delete:  stockUC.NewDeleteStockEntryUseCase(stockRepo),
		},
		Inventory: &InventoryHandlers{
			GetCurrentStock: inventoryUC.NewGetCurrentStockUseCase(inventoryRepo),
			GetAllStock:     inventoryUC.NewGetAllCurrentStockUseCase(inventoryRepo),
			GetSummary:      inventoryUC.NewGetInventorySummaryUseCase(inventoryRepo),
			GetLowStock:     inventoryUC.NewGetLowStockUseCase(inventoryRepo),
			GetOutOfStock:   inventoryUC.NewGetOutOfStockUseCase(inventoryRepo),
		},
	}

	router := setupRouter()
	registerRoutesProduct(router, handlers.Product)
	registerRoutesAuth(router, handlers.Auth)
	registerRoutesStock(router, handlers.Stock)
	registerRoutesInventory(router, handlers.Inventory)

	return &Application{
		Router:    router,
		Handlers:  handlers,
		DBCleanup: func() { db.Close() },
		DB:        db,
	}

}

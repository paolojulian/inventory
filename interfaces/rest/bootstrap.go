package rest

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"paolojulian.dev/inventory/config"
	"paolojulian.dev/inventory/infrastructure/postgres"
	productUC "paolojulian.dev/inventory/usecase/product_uc"
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

type Handlers struct {
	Product *ProductHandlers
	Auth    *AuthHandlers
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
	}

	router := setupRouter()
	registerRoutesProduct(router, handlers.Product)
	registerRoutesAuth(router, handlers.Auth)

	return &Application{
		Router:    router,
		Handlers:  handlers,
		DBCleanup: func() { db.Close() },
		DB:        db,
	}

}

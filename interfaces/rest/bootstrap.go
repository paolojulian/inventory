package rest

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"paolojulian.dev/inventory/config"
	"paolojulian.dev/inventory/infrastructure/postgres"
	productUC "paolojulian.dev/inventory/usecase/product"
)

type ProductHandlers struct {
	Create *productUC.CreateProductUseCase
}

type Handlers struct {
	Product *ProductHandlers
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

	postgres.MigrateSchema(db)

	// Wire the repo to use cases
	productRepo := postgres.NewProductRepository(db)
	handlers := &Handlers{
		Product: &ProductHandlers{
			Create: productUC.NewCreateProductUseCase(productRepo),
		},
	}

	router := gin.Default()
	registerRoutesProduct(router, handlers.Product)

	return &Application{
		Router:    router,
		Handlers:  handlers,
		DBCleanup: func() { db.Close() },
		DB:        db,
	}

}

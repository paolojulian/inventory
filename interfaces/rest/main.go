package rest

import (
	"log"

	"github.com/gin-gonic/gin"
	"paolojulian.dev/inventory/config"
	"paolojulian.dev/inventory/infrastructure/postgres"
	"paolojulian.dev/inventory/usecase/product"
)

func main() {
	// Load configs this early
	config.LoadConfig()

	db, err := postgres.NewPool()
	if err != nil {
		log.Fatalf("failed to connect to the database server. error: %v", err)
	}
	defer db.Close()

	// TODO: Transfer these wiring to another file
	productRepo := postgres.NewProductRepository(db)
	createProductUC := product.NewCreateProductUseCase(productRepo)

	// TODO: Transfer these to another file
	r := gin.Default()
	r.POST("/products", func(ctx *gin.Context) {
		// TODO: Parse input and call the usecase
		var input product.CreateProductInput

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(400, gin.H{
				"message": "Invalid Input",
			})
			return
		}

		result, err := createProductUC.Execute(ctx, input)
		if err != nil {
			status := 500
			if err == product.ErrSKUAlreadyExists {
				status = 409
			}
			ctx.JSON(status, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(201, gin.H{"product": result.ProductID})
	})

	r.Run()
}

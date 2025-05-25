package product_handler

import (
	"github.com/gin-gonic/gin"
	product_uc "paolojulian.dev/inventory/usecase/product"
)

func CreateHandler(uc *product_uc.CreateProductUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// TODO: Parse input and call the usecase
		var input product_uc.CreateProductInput

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(400, gin.H{
				"message": "Invalid Input",
			})
			return
		}

		result, err := uc.Execute(ctx, input)
		if err != nil {
			status := 500
			if err == product_uc.ErrSKUAlreadyExists {
				status = 409
			}
			ctx.JSON(status, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(201, gin.H{"product": result.Product})
	}
}

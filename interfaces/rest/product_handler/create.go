package product_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	product_uc "paolojulian.dev/inventory/usecase/product_uc"
)

func CreateHandler(uc *product_uc.CreateProductUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var input product_uc.CreateProductInput

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Input",
			})
			return
		}

		result, err := uc.Execute(ctx, input)
		if err != nil {
			status := http.StatusInternalServerError
			if err == product_uc.ErrSKUAlreadyExists {
				status = http.StatusConflict
			}
			ctx.JSON(status, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"product": result.Product})
	}
}

package product_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	product_uc "paolojulian.dev/inventory/usecase/product"
)

func DeactivateHandler(uc *product_uc.DeactivateProductUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productID := ctx.Param("id")
		if productID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Product ID is required.",
			})
			return
		}

		result, err := uc.Execute(ctx, productID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"product": result.Product})
	}
}

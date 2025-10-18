package inventory_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	inventory_uc "paolojulian.dev/inventory/usecase/inventory_uc"
)

func GetCurrentStockHandler(uc *inventory_uc.GetCurrentStockUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productID := ctx.Param("product_id")
		
		if productID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Product ID is required",
			})
			return
		}

		input := inventory_uc.GetCurrentStockInput{
			ProductID: productID,
		}

		output, err := uc.Execute(ctx, input)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Current stock retrieved successfully",
			"stock":   output.Stock,
		})
	}
}

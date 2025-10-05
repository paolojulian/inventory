package inventory_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	inventory_uc "paolojulian.dev/inventory/usecase/inventory_uc"
)

func GetOutOfStockHandler(uc *inventory_uc.GetOutOfStockUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		warehouseID := ctx.Query("warehouse_id")
		
		if warehouseID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Warehouse ID is required",
			})
			return
		}

		input := inventory_uc.GetOutOfStockInput{
			WarehouseID: warehouseID,
		}

		output, err := uc.Execute(ctx, input)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Out of stock products retrieved successfully",
			"stocks":  output.Stocks,
		})
	}
}

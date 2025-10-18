package inventory_handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	inventory_uc "paolojulian.dev/inventory/usecase/inventory_uc"
)

func GetLowStockHandler(uc *inventory_uc.GetLowStockUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		thresholdStr := ctx.DefaultQuery("threshold", "10")
		
		threshold, err := strconv.Atoi(thresholdStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid threshold value",
			})
			return
		}

		input := inventory_uc.GetLowStockInput{
			Threshold: threshold,
		}

		output, err := uc.Execute(ctx, input)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Low stock products retrieved successfully",
			"stocks":  output.Stocks,
		})
	}
}

package inventory_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	inventory_uc "paolojulian.dev/inventory/usecase/inventory_uc"
)

func GetAllStockHandler(uc *inventory_uc.GetAllCurrentStockUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		input := inventory_uc.GetAllCurrentStockInput{}

		output, err := uc.Execute(ctx, input)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "All stock retrieved successfully",
			"stocks":  output.Stocks,
		})
	}
}

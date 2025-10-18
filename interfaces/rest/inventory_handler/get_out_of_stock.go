package inventory_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	inventory_uc "paolojulian.dev/inventory/usecase/inventory_uc"
)

func GetOutOfStockHandler(uc *inventory_uc.GetOutOfStockUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		input := inventory_uc.GetOutOfStockInput{}

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

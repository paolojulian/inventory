package inventory_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	inventory_uc "paolojulian.dev/inventory/usecase/inventory_uc"
)

func GetAllStockHandler(uc *inventory_uc.GetAllCurrentStockUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input inventory_uc.GetAllCurrentStockInput

		if err := ctx.ShouldBindQuery(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Input",
			})
			return
		}

		// Do checks for pagination
		if !input.Pager.IsValid() {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid pager values",
			})
			return
		}

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
			"pager":   output.Pager,
		})
	}
}

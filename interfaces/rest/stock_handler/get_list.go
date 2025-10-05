package stock_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	stock_uc "paolojulian.dev/inventory/usecase/stock_uc"
)

func GetListHandler(uc *stock_uc.ListStockEntriesUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input stock_uc.ListStockEntriesInput

		if err := ctx.ShouldBindQuery(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Input",
			})
			return
		}

		output, err := uc.Execute(ctx, &input)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":       "Stock entries fetched successfully",
			"stock_entries": output.StockEntries,
			"total":         output.Total,
		})
		return
	}
}

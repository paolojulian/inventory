package stock_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	stock_uc "paolojulian.dev/inventory/usecase/stock_uc"
)

func DeleteHandler(uc *stock_uc.DeleteStockEntryUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		stockEntryID := ctx.Param("id")
		if stockEntryID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Stock entry ID is required",
			})
			return
		}

		input := stock_uc.DeleteStockEntryInput{
			StockEntryID: stockEntryID,
		}

		_, err := uc.Execute(ctx, input)
		if err != nil {
			status := http.StatusInternalServerError
			ctx.JSON(status, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Stock entry deleted successfully",
		})
	}
}

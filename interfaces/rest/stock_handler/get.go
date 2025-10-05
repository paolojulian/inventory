package stock_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	stock_uc "paolojulian.dev/inventory/usecase/stock_uc"
)

func GetHandler(uc *stock_uc.GetStockEntryUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		stockEntryID := ctx.Param("id")
		if stockEntryID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Stock entry ID is required",
			})
			return
		}

		output, err := uc.Execute(ctx, stockEntryID)
		if err != nil {
			status := http.StatusInternalServerError
			ctx.JSON(status, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":     "Stock entry fetched successfully",
			"stock_entry": output.StockEntry,
			"product":     output.Product,
			"warehouse":   output.Warehouse,
			"user":        output.User,
		})
	}
}

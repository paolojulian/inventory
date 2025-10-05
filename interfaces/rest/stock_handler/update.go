package stock_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	stock_uc "paolojulian.dev/inventory/usecase/stock_uc"
)

func UpdateHandler(uc *stock_uc.UpdateStockEntryUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		stockEntryID := ctx.Param("id")
		if stockEntryID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Stock entry ID is required",
			})
			return
		}

		var input stock_uc.UpdateStockEntryInput

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Input",
			})
			return
		}

		// Get user ID from context (set by auth middleware)
		userID, exists := ctx.Get("user_id")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "User not authenticated",
			})
			return
		}

		result, err := uc.Execute(ctx, stockEntryID, input, userID.(string))
		if err != nil {
			status := http.StatusInternalServerError
			if err == stock_uc.ErrInvalidReason {
				status = http.StatusBadRequest
			}
			ctx.JSON(status, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":     "Stock entry updated successfully",
			"stock_entry": result.StockEntry,
		})
	}
}

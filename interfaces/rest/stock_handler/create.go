package stock_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"paolojulian.dev/inventory/interfaces/rest/middleware"
	stock_uc "paolojulian.dev/inventory/usecase/stock_uc"
)

func CreateHandler(uc *stock_uc.CreateStockEntryUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input stock_uc.StockEntryInput

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Input",
			})
			return
		}

		// Get user ID from context (set by auth middleware)
		userID, exists := ctx.Get(middleware.UserIDKey)
		if !exists {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "User not authenticated",
			})
			return
		}

		result, err := uc.Execute(ctx, &input, userID.(string))
		if err != nil {
			status := http.StatusInternalServerError
			if err == stock_uc.ErrInvalidReason {
				status = http.StatusBadRequest
			}
			ctx.JSON(status, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"stock_entry": result.StockEntry})
	}
}

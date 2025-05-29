package product_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	productUC "paolojulian.dev/inventory/usecase/product_uc"
)

func UpdateHandler(uc *productUC.UpdateProductBasicUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productID := ctx.Param("id")
		if productID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Product ID is required",
			})
			return
		}

		var input productUC.UpdateProductBasicInput

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid input",
			})
			return
		}

		_, err := uc.Execute(ctx, productID, input)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Invalid input",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "product updated successfully",
		})
	}
}

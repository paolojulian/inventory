package product_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	productUC "paolojulian.dev/inventory/usecase/product"
)

func DeleteHandler(uc *productUC.DeleteProductUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input productUC.DeleteProductInput

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Input",
			})
			return
		}

		_, err := uc.Execute(ctx, input)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusNoContent, gin.H{
			"message": "Product successfully deleted",
		})
	}
}

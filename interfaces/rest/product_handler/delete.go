package product_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	productUC "paolojulian.dev/inventory/usecase/product"
)

func DeleteHandler(uc *productUC.DeleteProductUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productID := ctx.Param("id")
		if productID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Product ID is required",
			})
			return
		}

		_, err := uc.Execute(ctx, productUC.DeleteProductInput{ProductID: productID})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}

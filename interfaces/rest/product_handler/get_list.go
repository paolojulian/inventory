package product_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"paolojulian.dev/inventory/usecase/product_uc"
)

func GetListHandler(uc *product_uc.GetProductListUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input product_uc.GetProductListInput

		if err := ctx.ShouldBindQuery(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Input",
			})
			return
		}

		// Do checks for sorts
		if input.Sort != nil {
			if !input.Sort.Field.IsValid() {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "Invalid Sort Field",
				})
				return
			}

			if !input.Sort.Order.IsValid() {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "Invalid Sort Order",
				})
				return
			}
		}

		output, err := uc.Execute(ctx, input)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":  "Products fetched successfully",
			"products": output.Products,
			"pager":    output.Pager,
		})
		return
	}
}

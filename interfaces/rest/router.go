package rest

import (
	"github.com/gin-gonic/gin"
	"paolojulian.dev/inventory/interfaces/rest/product_handler"
)

func registerRoutesProduct(r *gin.Engine, handlers *ProductHandlers) {
	r.POST("/products", product_handler.CreateHandler(handlers.Create))
	r.DELETE("/products/:id", product_handler.DeleteHandler(handlers.Delete))
}

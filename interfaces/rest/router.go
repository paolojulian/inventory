package rest

import (
	"github.com/gin-gonic/gin"
	"paolojulian.dev/inventory/config"
	"paolojulian.dev/inventory/interfaces/rest/middleware"
	"paolojulian.dev/inventory/interfaces/rest/product_handler"
	middlewareTest "paolojulian.dev/inventory/tests/middleware"
)

func registerRoutesProduct(r *gin.Engine, handlers *ProductHandlers) {
	var productGroup *gin.RouterGroup
	if config.IsTestEnv() {
		productGroup = r.Group("/products", middlewareTest.TestAuthMiddleware())
	} else {
		productGroup = r.Group("/products", middleware.AuthMiddleware())
	}

	productGroup.POST("", product_handler.CreateHandler(handlers.Create))
	productGroup.DELETE("/:id", product_handler.DeleteHandler(handlers.Delete))
	productGroup.PUT("/:id", product_handler.UpdateHandler(handlers.Update))
	productGroup.POST("/:id/activate", product_handler.ActivateHandler(handlers.Activate))
	productGroup.POST("/:id/de-activate", product_handler.DeactivateHandler(handlers.Deactivate))
}

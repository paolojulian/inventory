package rest

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"paolojulian.dev/inventory/config"
	"paolojulian.dev/inventory/interfaces/rest/inventory_handler"
	"paolojulian.dev/inventory/interfaces/rest/middleware"
	"paolojulian.dev/inventory/interfaces/rest/product_handler"
	"paolojulian.dev/inventory/interfaces/rest/user_handler"
	middlewareTest "paolojulian.dev/inventory/tests/middleware"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	return router
}

func registerRoutesAuth(r *gin.Engine, handlers *AuthHandlers) {
	var authGroup *gin.RouterGroup
	if config.IsTestEnv() {
		authGroup = r.Group("/auth", middlewareTest.TestAuthMiddleware())
	} else {
		authGroup = r.Group("/auth", middleware.AuthMiddleware())
	}

	authGroup.POST("/login", user_handler.LoginHandler(handlers.Login))
	authGroup.POST("/logout", user_handler.LogoutHandler())
	authGroup.POST("/me", user_handler.MeHandler())
}

func registerRoutesProduct(r *gin.Engine, handlers *ProductHandlers) {
	var productGroup *gin.RouterGroup
	if config.IsTestEnv() {
		productGroup = r.Group("/products", middlewareTest.TestAuthMiddleware())
	} else {
		productGroup = r.Group("/products", middleware.AuthMiddleware())
	}

	productGroup.POST("", product_handler.CreateHandler(handlers.Create))
	productGroup.GET("", product_handler.GetListHandler(handlers.GetList))
	productGroup.DELETE("/:id", product_handler.DeleteHandler(handlers.Delete))
	productGroup.PUT("/:id", product_handler.UpdateHandler(handlers.Update))
	productGroup.POST("/:id/activate", product_handler.ActivateHandler(handlers.Activate))
	productGroup.POST("/:id/de-activate", product_handler.DeactivateHandler(handlers.Deactivate))
}

// func registerRoutesStock(r *gin.Engine, handlers *StockHandlers) {
// 	var stockGroup *gin.RouterGroup
// 	if config.IsTestEnv() {
// 		stockGroup = r.Group("/stock-entries", middlewareTest.TestAuthMiddleware())
// 	} else {
// 		stockGroup = r.Group("/stock-entries", middleware.AuthMiddleware())
// 	}

// 	stockGroup.POST("", stock_handler.CreateHandler(handlers.Create))
// 	stockGroup.GET("", stock_handler.GetListHandler(handlers.GetList))
// 	stockGroup.GET("/:id", stock_handler.GetHandler(handlers.Get))
// 	stockGroup.PUT("/:id", stock_handler.UpdateHandler(handlers.Update))
// 	stockGroup.DELETE("/:id", stock_handler.DeleteHandler(handlers.Delete))
// }

func registerRoutesInventory(r *gin.Engine, handlers *InventoryHandlers) {
	var inventoryGroup *gin.RouterGroup
	if config.IsTestEnv() {
		inventoryGroup = r.Group("/inventory", middlewareTest.TestAuthMiddleware())
	} else {
		inventoryGroup = r.Group("/inventory", middleware.AuthMiddleware())
	}

	inventoryGroup.GET("/all-stock", inventory_handler.GetAllStockHandler(handlers.GetAllStock))
	inventoryGroup.GET("/current-stock/:product_id", inventory_handler.GetCurrentStockHandler(handlers.GetCurrentStock))
	inventoryGroup.GET("/summary", inventory_handler.GetInventorySummaryHandler(handlers.GetSummary))
	inventoryGroup.GET("/low-stock", inventory_handler.GetLowStockHandler(handlers.GetLowStock))
	inventoryGroup.GET("/out-of-stock", inventory_handler.GetOutOfStockHandler(handlers.GetOutOfStock))
}

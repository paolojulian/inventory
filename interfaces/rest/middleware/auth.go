package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"paolojulian.dev/inventory/infrastructure/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr, err := ctx.Cookie("access_token")
		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		userID, err := auth.ParseToken(tokenStr)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Store userID in Gin context for later use in handlers
		ctx.Set("userID", userID)

		ctx.Next()
	}
}
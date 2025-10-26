package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"paolojulian.dev/inventory/infrastructure/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the Authorization header
        authHeader := ctx.GetHeader("Authorization")
        if authHeader == "" {
            log.Printf("AuthMiddleware: missing Authorization header")
            ctx.AbortWithStatus(http.StatusUnauthorized)
            return
        }

		// Extract the token from "Bearer <token>"
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            log.Printf("AuthMiddleware: invalid auth header format: %q", authHeader)
            ctx.AbortWithStatus(http.StatusUnauthorized)
            return
        }

		tokenStr := parts[1]

        userID, err := auth.ParseToken(tokenStr)
        if err != nil {
            log.Printf("AuthMiddleware: token parse error: %v", err)
            ctx.AbortWithStatus(http.StatusUnauthorized)
            return
        }

		// Store userID in Gin context for later use in handlers
		ctx.Set("userID", userID)

		ctx.Next()
	}
}
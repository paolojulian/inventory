package user_handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"paolojulian.dev/inventory/infrastructure/auth"
)

// Checks if the user is logged in
func MeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the Authorization header
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		// Extract the token from "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Printf("token is not valid: %v")
			ctx.Status(http.StatusUnauthorized)
			return
		}

		token := parts[1]

		if !auth.IsTokenValid(token) {
			log.Printf("token is not valid: %v", token)
			ctx.Status(http.StatusUnauthorized)
			return
		}

		ctx.Status(http.StatusOK)
	}
}

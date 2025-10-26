package user_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// With JWT token auth, logout is handled client-side by removing the token
		// Server doesn't need to do anything special
		ctx.JSON(http.StatusOK, gin.H{
			"message": "logged out successfully.",
		})
	}
}

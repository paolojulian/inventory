package user_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"paolojulian.dev/inventory/infrastructure/auth"
)

// Checks if the user is logged in
func MeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie(accessTokenKey)
		if err != nil {
			removeCookie(ctx)
			ctx.Status(http.StatusUnauthorized)
			return
		}

		if !auth.IsTokenValid(token) {
			removeCookie(ctx)
			ctx.Status(http.StatusUnauthorized)
			return
		}

		ctx.Status(http.StatusOK)
	}
}

func removeCookie(ctx *gin.Context) {
	// Remove cookie
	ctx.SetCookie(accessTokenKey, "", -1, "/", "", true, true)
}

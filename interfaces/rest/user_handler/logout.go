package user_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Remove cookie
		ctx.SetCookie(accessTokenKey, "", -1, "/", "", true, true)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "logged out successfully.",
		})
	}
}

package user_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"paolojulian.dev/inventory/usecase/user_uc"
)

func LogoutHandler(uc *user_uc.LoginUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Remove cookie
		ctx.SetCookie("access_token", "", -1, "/", "", true, true)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "logged out successfully.",
		})
	}
}

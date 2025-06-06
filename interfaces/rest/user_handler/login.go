package user_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"paolojulian.dev/inventory/usecase/user_uc"
)

func LoginHandler(uc *user_uc.LoginUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input user_uc.LoginInput
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		result, err := uc.Execute(ctx, &input)
		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			return
		}

		ctx.SetCookie("access_token", string(result.Token), 86400, "/", "", true, true) // 1-day cookie, Secure + HTTPOnly

		ctx.JSON(http.StatusOK, gin.H{
			"message": "logged in successfully.",
		})
	}
}

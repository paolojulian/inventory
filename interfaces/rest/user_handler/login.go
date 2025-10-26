package user_handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"paolojulian.dev/inventory/usecase/user_uc"
)

func LoginHandler(uc *user_uc.LoginUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input user_uc.LoginInput
		if err := ctx.ShouldBindJSON(&input); err != nil {
			log.Printf("error binding JSON: %v", err)
			ctx.Status(http.StatusUnauthorized)
			return
		}

		result, err := uc.Execute(ctx, &input)
		if err != nil {
			log.Printf("error logging in: %v", err)
			ctx.Status(http.StatusUnauthorized)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "logged in successfully.",
			"token":   string(result.Token),
		})
	}
}

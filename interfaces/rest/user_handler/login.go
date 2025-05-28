package user_handler

import (
	"github.com/gin-gonic/gin"
	"paolojulian.dev/inventory/usecase/user_uc"
)

func LoginHandler(uc *user_uc.LoginUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

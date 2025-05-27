package middleware

import "github.com/gin-gonic/gin"

var (
	FakeUserID = "test-user-id"
)

func TestAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("userID", FakeUserID) // fake user
		ctx.Next()
	}
}

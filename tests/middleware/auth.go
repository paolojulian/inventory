package middleware

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"paolojulian.dev/inventory/infrastructure/postgres"
)

const (
	FakeUserID = "550e8400-e29b-41d4-a716-446655440001"
)

func TestAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Try to find admin user from database
		db, err := postgres.NewPool()
		if err != nil {
			log.Printf("TestAuthMiddleware: failed to connect to database: %v", err)
			ctx.Set("userID", FakeUserID)
			ctx.Next()
			return
		}
		defer db.Close()

		userRepo := postgres.NewUserRepository(db)
		adminUser, err := userRepo.FindByUsername(context.Background(), "admin")
		if err != nil || adminUser == nil {
			log.Printf("TestAuthMiddleware: failed to find admin user: %v", err)
			ctx.Set("userID", FakeUserID)
			ctx.Next()
			return
		}

		ctx.Set("userID", adminUser.ID)
		ctx.Next()
	}
}

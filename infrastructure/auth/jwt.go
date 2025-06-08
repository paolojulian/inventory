package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY")) // 🔐 Use a secure key from env in prod
type AccessToken string

func NewAccessToken(userID string) (AccessToken, error) {
	claims := jwt.MapClaims{
		"sub": userID,                                // subject
		"exp": time.Now().Add(24 * time.Hour).Unix(), // expires in 24h
		"iat": time.Now().Unix(),                     // issued at
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString(jwtSecret)

	return AccessToken(accessToken), err
}

func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	claims := token.Claims.(jwt.MapClaims)
	sub, ok := claims["sub"].(string)
	if !ok {
		return "", jwt.ErrTokenMalformed
	}
	return sub, nil
}

func IsTokenValid(tokenStr string) bool {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	return err == nil && token.Valid
}

package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Todo : Implementing JWT Validation logic

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authHeader := ctx.Request.Header.Get("Authorization")

		if authHeader == "" {
			ctx.JSON(401, gin.H{"error": "Authorization header is required"})
			ctx.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			ctx.JSON(401, gin.H{"error": "Authorization header is invalid"})
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(401, gin.H{"error": "Authorization header is invalid"})
			ctx.Abort()
			return
		}

		ctx.Next()

	}
}

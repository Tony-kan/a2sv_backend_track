package middleware

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Todo : Implementing JWT Validation logic

// var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
var jwtSecret = []byte("secret")

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

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(401, gin.H{"error": "Invalid token claims"})
			ctx.Abort()
			return
		}

		roleStr, ok := claims["role"].(string)

		if !ok {
			ctx.JSON(401, gin.H{"error": "Role claim is missing or invalid"})
			ctx.Abort()
			return
		}

		ctx.Set("role", roleStr)
		ctx.Next()

	}
}

func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, exists := ctx.Get("role")
		if !exists {
			ctx.JSON(403, gin.H{"error": "Forbidden: No role associated with user"})
			ctx.Abort()
			return
		}

		roleStr, ok := role.(string)
		if !ok {
			ctx.JSON(403, gin.H{"error": "Forbidden: Invalid role type"})
			ctx.Abort()
			return
		}

		for _, allowedRole := range allowedRoles {
			if roleStr == allowedRole {
				ctx.Next()
				return
			}
		}

		ctx.JSON(403, gin.H{"error": "Forbidden: Insufficient permissions"})
		ctx.Abort()
	}
}

package infrastructure

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Todo : functions to generate and validate JWT tokens

var jwtSecret = []byte("secret")

func GenerateJWTToken(userID string, email, role string) (string, error) {
	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"exp":     jwt.TimeFunc().Add(72 * time.Hour).Unix(),
	})

	// Sign the token with the secret key
	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func ValidateJWTToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("token validation failed: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// return nil, fmt.Errorf("invalid token claims")
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims format")
	}

	// Validate required claims
	requiredClaims := []string{"user_id", "email", "role", "exp"}
	for _, claim := range requiredClaims {
		if _, ok := claims[claim]; !ok {
			return nil, fmt.Errorf("missing required claim: %s", claim)
		}
	}

	if _, ok := claims["role"].(string); !ok {
		return nil, fmt.Errorf("invalid role type")
	}

	// Validate expiration dateee
	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid expiration format")
	}

	expTime := time.Unix(int64(exp), 0)
	if time.Now().After(expTime) {
		return nil, fmt.Errorf("token expired")
	}

	return claims, nil
}

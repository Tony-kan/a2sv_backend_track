package infrastructure_test

import (
	"testing"
	"time"

	"task_8_task_management_api_testing/infrastructure"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWTToken_Success(t *testing.T) {
	token, err := infrastructure.GenerateJWTToken("123", "test@example.com", "admin")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedToken, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	claims := parsedToken.Claims.(jwt.MapClaims)
	assert.Equal(t, "123", claims["user_id"])
	assert.Equal(t, "test@example.com", claims["email"])
	assert.Equal(t, "admin", claims["role"])
}

func TestValidateJWTToken_Success(t *testing.T) {
	token, _ := infrastructure.GenerateJWTToken("123", "test@example.com", "admin")
	claims, err := infrastructure.ValidateJWTToken(token)
	assert.NoError(t, err)
	assert.Equal(t, "123", claims["user_id"])
}

func TestValidateJWTToken_InvalidSignature(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": "123",
		"email":   "test@example.com",
		"role":    "admin",
		"exp":     time.Now().Add(time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString([]byte("wrong_secret"))

	_, err := infrastructure.ValidateJWTToken(tokenString)
	assert.ErrorContains(t, err, "token validation failed")
}

func TestValidateJWTToken_ExpiredToken(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": "123",
		"email":   "test@example.com",
		"role":    "admin",
		"exp":     time.Now().Add(-1 * time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString([]byte("secret"))

	_, err := infrastructure.ValidateJWTToken(tokenString)
	assert.ErrorContains(t, err, "is expired")
}

func TestValidateJWTToken_MissingClaims(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": "123",
		"exp":     time.Now().Add(time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString([]byte("secret"))

	_, err := infrastructure.ValidateJWTToken(tokenString)
	assert.ErrorContains(t, err, "missing required claim: email")
}

package infrastructure_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"task_8_task_management_api_testing/infrastructure"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware_MissingAuthorizationHeader(t *testing.T) {
	router := gin.New()
	router.Use(infrastructure.AuthMiddleware())
	router.GET("/test", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Authorization header is required")
}

func TestAuthMiddleware_InvalidHeaderFormat(t *testing.T) {
	router := gin.New()
	router.Use(infrastructure.AuthMiddleware())
	router.GET("/test", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "InvalidFormat")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Authorization header is invalid")
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	router := gin.New()
	router.Use(infrastructure.AuthMiddleware())
	router.GET("/test", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	// Generate token with invalid signature
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": "123",
		"exp":     time.Now().Add(time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString([]byte("wrong_secret"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Authorization header is invalid")
}

func TestAuthMiddleware_ValidToken(t *testing.T) {
	userID := "123"
	email := "test@example.com"
	role := "admin"
	token, err := infrastructure.GenerateJWTToken(userID, email, role)
	assert.NoError(t, err)

	router := gin.New()
	router.Use(infrastructure.AuthMiddleware())
	router.GET("/test", func(ctx *gin.Context) {
		assert.Equal(t, userID, ctx.GetString("user_id"))
		assert.Equal(t, email, ctx.GetString("email"))
		assert.Equal(t, role, ctx.GetString("role"))
		ctx.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRequireRole_AllowedRole(t *testing.T) {
	token, _ := infrastructure.GenerateJWTToken("123", "test@example.com", "admin")

	router := gin.New()
	router.Use(infrastructure.AuthMiddleware())
	router.GET("/test", infrastructure.RequireRole("admin"), func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRequireRole_ForbiddenRole(t *testing.T) {
	token, _ := infrastructure.GenerateJWTToken("123", "test@example.com", "user")

	router := gin.New()
	router.Use(infrastructure.AuthMiddleware())
	router.GET("/test", infrastructure.RequireRole("admin"), func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Contains(t, w.Body.String(), "Insufficient permissions")
	// assert.Contains(t, w.Body.String(), "Insufficient permissions")
    assert.Contains(t, w.Body.String(), "Role 'user'")
    assert.Contains(t, w.Body.String(), "Allowed roles: [admin]")
}

func TestRequireRole_MissingRole(t *testing.T) {
	router := gin.New()
	router.GET("/test", infrastructure.RequireRole("admin"), func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Contains(t, w.Body.String(), "No role associated with user")
}

func TestRequireRole_InvalidRoleType(t *testing.T) {
	router := gin.New()
	router.Use(func(ctx *gin.Context) {
		ctx.Set("role", 123) // Invalid type
		ctx.Next()
	})
	router.GET("/test", infrastructure.RequireRole("admin"), func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid role type")
}

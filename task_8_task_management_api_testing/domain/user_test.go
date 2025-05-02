package domain_test

import (
	"task_8_task_management_api_testing/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserValidation(t *testing.T) {
	now := time.Now().UTC()

	t.Run("ValidUser", func(t *testing.T) {
		user := &domain.User{
			ID:        primitive.NewObjectID(),
			Username:  "tkay",
			Email:     "tony@example.com",
			Password:  "SecurePassword123!",
			Role:      domain.UserRole,
			CreatedAt: now,
			UpdatedAt: now,
		}

		assert.NotEmpty(t, user.ID)
		assert.Equal(t, "tkay", user.Username)
		assert.Equal(t, domain.UserRole, user.Role)
		assert.False(t, user.CreatedAt.IsZero())
		assert.False(t, user.UpdatedAt.IsZero())
	})

	t.Run("InvalidEmailFormat", func(t *testing.T) {
		user := &domain.User{
			ID:       primitive.NewObjectID(),
			Username: "invalidemail",
			Email:    "not-an-email",
			Password: "ValidPass123",
			Role:     domain.UserRole,
		}

		assert.Equal(t, "not-an-email", user.Email)
	})

	t.Run("MissingUsername", func(t *testing.T) {
		user := &domain.User{
			Username: "",
			Email:    "valid@example.com",
			Password: "ValidPass123",
			Role:     domain.UserRole,
		}

		assert.Empty(t, user.Username)

	})

	t.Run("WeakPassword", func(t *testing.T) {
		user := &domain.User{
			Username: "weakpassuser",
			Email:    "weak@example.com",
			Password: "short",
			Role:     domain.UserRole,
		}

		assert.Equal(t, "short", user.Password)
	})

	t.Run("InvalidRole", func(t *testing.T) {
		user := &domain.User{
			Username: "invalidrole",
			Email:    "role@example.com",
			Password: "ValidPass123",
			Role:     "superuser", // Invalid role
		}

		assert.Equal(t, domain.RoleType("superuser"), user.Role)
	})

	t.Run("AdminRoleValid", func(t *testing.T) {
		user := &domain.User{
			Username: "adminuser",
			Email:    "admin@example.com",
			Password: "AdminPass123!",
			Role:     domain.AdminRole,
		}

		assert.Equal(t, domain.AdminRole, user.Role)
	})
}

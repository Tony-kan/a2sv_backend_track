package infrastructure_test

import (
	"testing"

	"task_8_task_management_api_testing/infrastructure"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword_Success(t *testing.T) {
	hashed, err := infrastructure.HashPassword("password123")
	assert.NoError(t, err)
	assert.NotEmpty(t, hashed)
}

func TestComparePassword_Correct(t *testing.T) {
	hashed, _ := infrastructure.HashPassword("password123")
	err := infrastructure.ComparePassword(hashed, "password123")
	assert.NoError(t, err)
}

func TestComparePassword_Incorrect(t *testing.T) {
	hashed, _ := infrastructure.HashPassword("password123")
	err := infrastructure.ComparePassword(hashed, "wrongpassword")
	assert.Error(t, err)
}

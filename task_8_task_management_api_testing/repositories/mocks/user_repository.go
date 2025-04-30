package mocks

import (
	"context"
	"task_8_task_management_api_testing/domain"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) RegisterUser(ctx context.Context, user *domain.User) (string, error) {
	args := m.Called(ctx, user)
	return args.String(0), args.Error(1)
}

func (m *MockUserRepository) LoginUser(ctx context.Context, loginRequest domain.LoginRequest) (string, error) {
	args := m.Called(ctx, loginRequest)
	return args.String(0), args.Error(1)
}

func (m *MockUserRepository) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*domain.User), args.Error(1)
}
package usecases_test

import (
	"context"
	"task_8_task_management_api_testing/domain"
	"task_8_task_management_api_testing/repositories/mocks"
	"task_8_task_management_api_testing/usecases"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserUsecase_RegisterUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		usecase := usecases.NewUserUsecase(mockRepo, 2*time.Second)

		user := &domain.User{
			Username: "testuser",
			Email:    "test@example.com",
			Password: "securepassword",
		}
		expectedID := "user-123"

		mockRepo.On("RegisterUser", mock.Anything, user).Return(expectedID, nil)

		id, err := usecase.RegisterUser(context.Background(), user)

		assert.NoError(t, err)
		assert.Equal(t, expectedID, id)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DuplicateEmail", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		usecase := usecases.NewUserUsecase(mockRepo, 2*time.Second)

		user := &domain.User{Email: "exists@example.com"}
		mockRepo.On("RegisterUser", mock.Anything, user).Return("", domain.ErrUserExists)

		_, err := usecase.RegisterUser(context.Background(), user)

		assert.ErrorIs(t, err, domain.ErrUserExists)
		mockRepo.AssertExpectations(t)
	})
}

func TestUserUsecase_LoginUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		usecase := usecases.NewUserUsecase(mockRepo, 2*time.Second)

		loginReq := domain.LoginRequest{
			Email:    "valid@example.com",
			Password: "correct-password",
		}
		expectedToken := "jwt-token-123"

		mockRepo.On("LoginUser", mock.Anything, loginReq).Return(expectedToken, nil)

		token, err := usecase.LoginUser(context.Background(), loginReq)

		assert.NoError(t, err)
		assert.Equal(t, expectedToken, token)
		mockRepo.AssertExpectations(t)
	})

	t.Run("InvalidCredentials", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		usecase := usecases.NewUserUsecase(mockRepo, 2*time.Second)

		loginReq := domain.LoginRequest{
			Email:    "invalid@example.com",
			Password: "wrong-password",
		}
		mockRepo.On("LoginUser", mock.Anything, loginReq).Return("", domain.ErrInvalidCredentials)

		_, err := usecase.LoginUser(context.Background(), loginReq)

		assert.ErrorIs(t, err, domain.ErrInvalidCredentials)
		mockRepo.AssertExpectations(t)
	})
}

func TestUserUsecase_GetAllUsers(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		usecase := usecases.NewUserUsecase(mockRepo, 2*time.Second)

		id1, _ := primitive.ObjectIDFromHex("63c0a4e72772c6a9e45678b1")
		id2, _ := primitive.ObjectIDFromHex("63c0a4e72772c6a9e45678b2")

		expectedUsers := []*domain.User{
			{ID: id1, Username: "user1"},
			{ID: id2, Username: "user2"},
		}
		mockRepo.On("GetAllUsers", mock.Anything).Return(expectedUsers, nil)

		users, err := usecase.GetAllUsers(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, expectedUsers, users)
		mockRepo.AssertExpectations(t)
	})

	t.Run("EmptyList", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		usecase := usecases.NewUserUsecase(mockRepo, 2*time.Second)

		mockRepo.On("GetAllUsers", mock.Anything).Return([]*domain.User{}, nil)

		users, err := usecase.GetAllUsers(context.Background())

		assert.NoError(t, err)
		assert.Empty(t, users)
		mockRepo.AssertExpectations(t)
	})

	t.Run("AccessDenied", func(t *testing.T) {
		mockRepo := new(mocks.MockUserRepository)
		usecase := usecases.NewUserUsecase(mockRepo, 2*time.Second)

		// Return nil slice with error
		mockRepo.On("GetAllUsers", mock.Anything).Return(([]*domain.User)(nil), domain.ErrAccessDenied)

		_, err := usecase.GetAllUsers(context.Background())

		assert.ErrorIs(t, err, domain.ErrAccessDenied)
		mockRepo.AssertExpectations(t)
	})
}

package usecases

import (
	"context"
	"task_7_task_management_api_refactoring/domain"
	"time"
)

type UserUsecase struct {
	UserRepository domain.UserRepository
	ContextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &UserUsecase{
		UserRepository: userRepository,
		ContextTimeout: timeout,
	}
}

func (usecase *UserUsecase) RegisterUser(user *domain.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), usecase.ContextTimeout)
	defer cancel()

	return usecase.UserRepository.RegisterUser(ctx, user)
}
func (usecase *UserUsecase) LoginUser(loginRequest domain.LoginRequest) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), usecase.ContextTimeout)
	defer cancel()

	return usecase.UserRepository.LoginUser(ctx, loginRequest)
}
func (usecase *UserUsecase) GetAllUsers() ([]*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), usecase.ContextTimeout)
	defer cancel()

	return usecase.UserRepository.GetAllUsers(ctx)
}

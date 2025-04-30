package mocks

import (
	"context"
	"task_8_task_management_api_testing/domain"

	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) AddTask(ctx context.Context, task *domain.Task) (string, error) {
	args := m.Called(ctx, task)
	return args.String(0), args.Error(1)
}

func (m *MockTaskRepository) RemoveTask(ctx context.Context, taskID string) error {
	args := m.Called(ctx, taskID)
	return args.Error(0)
}

func (m *MockTaskRepository) GetTaskById(ctx context.Context, taskID string) (*domain.Task, error) {
	args := m.Called(ctx, taskID)
	return args.Get(0).(*domain.Task), args.Error(1)
}

func (m *MockTaskRepository) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*domain.Task), args.Error(1)
}

func (m *MockTaskRepository) UpdateTask(ctx context.Context, taskID string, updateFields map[string]interface{}) error {
	args := m.Called(ctx, taskID, updateFields)
	return args.Error(0)
}

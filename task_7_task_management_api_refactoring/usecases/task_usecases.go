package usecases

import (
	"context"
	"task_7_task_management_api_refactoring/domain"
	"time"
)

type TaskUsecase struct {
	TaskRepository domain.TaskRepository
	ContextTimeout time.Duration
}

func NewTaskUsecase(taskRepository domain.TaskRepository, timeout time.Duration) domain.TaskUsecase {
	return &TaskUsecase{
		TaskRepository: taskRepository,
		ContextTimeout: timeout,
	}
}
func (usecase *TaskUsecase) AddTask(ctx context.Context, task *domain.Task) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.ContextTimeout)
	defer cancel()
	return usecase.TaskRepository.AddTask(ctx, task)
}

func (usecase *TaskUsecase) RemoveTask(ctx context.Context, taskID string) error {
	ctx, cancel := context.WithTimeout(ctx, usecase.ContextTimeout)
	defer cancel()
	return usecase.TaskRepository.RemoveTask(ctx, taskID)
}
func (usecase *TaskUsecase) GetTaskById(ctx context.Context, taskID string) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.ContextTimeout)
	defer cancel()
	return usecase.TaskRepository.GetTaskById(ctx, taskID)
}
func (usecase *TaskUsecase) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.ContextTimeout)
	defer cancel()
	return usecase.TaskRepository.GetAllTasks(ctx)
}

func (usecase *TaskUsecase) UpdateTask(ctx context.Context, updateFields map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, usecase.ContextTimeout)
	defer cancel()
	return usecase.TaskRepository.UpdateTask(ctx, updateFields)
}

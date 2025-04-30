package usecases_test

import (
	"context"
	"errors"
	"task_8_task_management_api_testing/domain"
	"task_8_task_management_api_testing/repositories/mocks"
	"task_8_task_management_api_testing/usecases"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTaskUsecase_AddTask(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.MockTaskRepository)
		usecase := usecases.NewTaskUsecase(mockRepo, 2*time.Second)

		task := &domain.Task{Title: "Test Task"}
		expectedID := "task-123"

		mockRepo.On("AddTask", mock.Anything, task).Return(expectedID, nil)

		id, err := usecase.AddTask(context.Background(), task)

		assert.NoError(t, err)
		assert.Equal(t, expectedID, id)
		mockRepo.AssertExpectations(t)
	})

	t.Run("RepositoryError", func(t *testing.T) {
		mockRepo := new(mocks.MockTaskRepository)
		usecase := usecases.NewTaskUsecase(mockRepo, 2*time.Second)

		task := &domain.Task{Title: "Test Task"}
		expectedErr := errors.New("repository error")

		mockRepo.On("AddTask", mock.Anything, task).Return("", expectedErr)

		_, err := usecase.AddTask(context.Background(), task)

		assert.ErrorIs(t, err, expectedErr)
		mockRepo.AssertExpectations(t)
	})

	// t.Run("ContextTimeout", func(t *testing.T) {
	// 	mockRepo := new(mocks.MockTaskRepository)
	// 	usecase := usecases.NewTaskUsecase(mockRepo, time.Nanosecond)

	// 	ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
	// 	defer cancel()

	// 	_, err := usecase.AddTask(ctx, &domain.Task{})

	// 	assert.ErrorContains(t, err, "context deadline exceeded")
	// })
	t.Run("ContextTimeout", func(t *testing.T) {
		mockRepo := new(mocks.MockTaskRepository)
		usecase := usecases.NewTaskUsecase(mockRepo, time.Nanosecond)

		ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
		defer cancel()

		// Don't set up any mock expectations
		// Sleep to ensure timeout occurs before any repository call
		time.Sleep(1 * time.Millisecond)

		_, err := usecase.AddTask(ctx, &domain.Task{
			Title:   "Test Task",
			DueDate: time.Now().AddDate(0, 0, 1),
			Status:  "Pending",
		})

		assert.ErrorContains(t, err, "context deadline exceeded")
		mockRepo.AssertNotCalled(t, "AddTask") // Verify no repository call
	})
}

func TestTaskUsecase_GetTaskById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.MockTaskRepository)
		usecase := usecases.NewTaskUsecase(mockRepo, 2*time.Second)
		id1, _ := primitive.ObjectIDFromHex("63c0a4e72772c6a9e45678b6")

		expectedTask := &domain.Task{ID: id1, Title: "Test Task"}
		mockRepo.On("GetTaskById", mock.Anything, "task-123").Return(expectedTask, nil)

		task, err := usecase.GetTaskById(context.Background(), "task-123")

		assert.NoError(t, err)
		assert.Equal(t, expectedTask, task)
		mockRepo.AssertExpectations(t)
	})

	t.Run("NotFound", func(t *testing.T) {
		mockRepo := new(mocks.MockTaskRepository)
		usecase := usecases.NewTaskUsecase(mockRepo, 2*time.Second)

		mockRepo.On("GetTaskById", mock.Anything, "invalid-id").Return((*domain.Task)(nil), errors.New("not found"))

		_, err := usecase.GetTaskById(context.Background(), "invalid-id")

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestTaskUsecase_UpdateTask(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.MockTaskRepository)
		usecase := usecases.NewTaskUsecase(mockRepo, 2*time.Second)

		updateFields := map[string]interface{}{"status": "completed"}
		mockRepo.On("UpdateTask", mock.Anything, "task-123", updateFields).Return(nil)

		err := usecase.UpdateTask(context.Background(), "task-123", updateFields)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("InvalidFields", func(t *testing.T) {
		mockRepo := new(mocks.MockTaskRepository)
		usecase := usecases.NewTaskUsecase(mockRepo, 2*time.Second)

		updateFields := map[string]interface{}{"invalid_field": "value"}
		mockRepo.On("UpdateTask", mock.Anything, "task-123", updateFields).Return(errors.New("invalid field"))

		err := usecase.UpdateTask(context.Background(), "task-123", updateFields)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestTaskUsecase_RemoveTask(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.MockTaskRepository)
		usecase := usecases.NewTaskUsecase(mockRepo, 2*time.Second)

		mockRepo.On("RemoveTask", mock.Anything, "task-123").Return(nil)

		err := usecase.RemoveTask(context.Background(), "task-123")

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("NotFound", func(t *testing.T) {
		mockRepo := new(mocks.MockTaskRepository)
		usecase := usecases.NewTaskUsecase(mockRepo, 2*time.Second)

		mockRepo.On("RemoveTask", mock.Anything, "invalid-id").Return(errors.New("not found"))

		err := usecase.RemoveTask(context.Background(), "invalid-id")

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestTaskUsecase_GetAllTasks(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockRepo := new(mocks.MockTaskRepository)
		usecase := usecases.NewTaskUsecase(mockRepo, 2*time.Second)

		id1, _ := primitive.ObjectIDFromHex("63c0a4e72772c6a9e45678b3")
		id2, _ := primitive.ObjectIDFromHex("63c0a4e72772c6a9e45678b4")

		expectedTasks := []*domain.Task{
			{ID: id1, Title: "Task 1"},
			{ID: id2, Title: "Task 2"},
		}
		mockRepo.On("GetAllTasks", mock.Anything).Return(expectedTasks, nil)

		tasks, err := usecase.GetAllTasks(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, expectedTasks, tasks)
		mockRepo.AssertExpectations(t)
	})

	t.Run("EmptyResult", func(t *testing.T) {
		mockRepo := new(mocks.MockTaskRepository)
		usecase := usecases.NewTaskUsecase(mockRepo, 2*time.Second)

		mockRepo.On("GetAllTasks", mock.Anything).Return([]*domain.Task{}, nil)

		tasks, err := usecase.GetAllTasks(context.Background())

		assert.NoError(t, err)
		assert.Empty(t, tasks)
		mockRepo.AssertExpectations(t)
	})
}

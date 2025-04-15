package services

import (
	"errors"
	"fmt"
	"task_4_task_management_api/models"
	"time"
)

// Todo : create an interface,services and constructor for the task service
// Todo : implement the methods of the interface
// Todo : create error handling methods

var (
	ErrTaskNotFound  = errors.New("task not found")
	ErrTaskExists    = errors.New("task already exists")
	ErrInvalidTask   = errors.New("invalid task data")
	ErrInvalidTaskID = errors.New("invalid task ID")
)

type TaskServices interface {
	AddTask(task models.Task) error
	RemoveTask(taskID string) error
	GetTaskById(taskID string) (models.Task, error)
	GetAllTasks() ([]models.Task, error)
	UpdateTask(taskID string, task models.Task) error
}

type TaskService struct {
	tasks map[string]models.Task
}

// Constructor which initializes 2 tasks in the memery during startup
func NewTaskService() TaskServices {
	service := &TaskService{
		tasks: make(map[string]models.Task),
	}
	defaultTask := []models.Task{
		{
			ID:          "1",
			Title:       "Task 1",
			Description: "Description for Task 1",
			Status:      "Pending",
			DueDate:     time.Now(),
		},
		{
			ID:          "2",
			Title:       "Task 2",
			Description: "Description for Task 2",
			Status:      "Completed",
			DueDate:     time.Now().AddDate(0, 0, 1),
		},
	}
	for _, task := range defaultTask {
		service.tasks[task.ID] = task
	}

	return service
}

func (service *TaskService) AddTask(task models.Task) error {
	for _, tsk := range service.tasks {
		if task.ID == tsk.ID {
			return fmt.Errorf("task with id %s already exists", task.ID)
		}
	}

	if task.ID == "" || task.Title == "" {
		return ErrInvalidTask
	}
	task.Status = "Pending"

	service.tasks[task.ID] = task
	return nil
}

func (service *TaskService) RemoveTask(taskID string) error {
	for _, task := range service.tasks {
		if task.ID == taskID {
			delete(service.tasks, taskID)
			return nil
		}
	}
	return fmt.Errorf("task with id %s not found", taskID)
}

func (service *TaskService) GetTaskById(taskID string) (models.Task, error) {
	for _, task := range service.tasks {
		if task.ID == taskID {
			return task, nil
		}
	}
	return models.Task{}, fmt.Errorf("task with id %s not found", taskID)
}

func (service *TaskService) GetAllTasks() ([]models.Task, error) {
	if len(service.tasks) == 0 {
		return []models.Task{}, nil
	}
	tasks := make([]models.Task, 0, len(service.tasks))
	for _, task := range service.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (service *TaskService) UpdateTask(taskID string, task models.Task) error {
	for i, tsk := range service.tasks {
		if tsk.ID == taskID {
			// Update the task on the specified fields
			taskToUpdate := service.tasks[i]
			if task.Title != "" {
				taskToUpdate.Title = task.Title
			}
			if task.Description != "" {
				taskToUpdate.Description = task.Description
			}
			if task.Status != "" {
				taskToUpdate.Status = task.Status
			}
			service.tasks[i] = taskToUpdate
			return nil
		}
	}
	return fmt.Errorf("task with id %s not found", taskID)
}

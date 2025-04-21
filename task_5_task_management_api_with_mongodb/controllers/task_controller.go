package controllers

import (
	services "task_5_task_management_api_with_mongodb/data"
	"task_5_task_management_api_with_mongodb/models"
)

type TaskController struct {
	taskService services.TaskServices
}

func NewTaskController(service services.TaskServices) *TaskController {
	return &TaskController{
		taskService: service,
	}
}

func (controller *TaskController) AddTask(task models.Task) error {
	return controller.taskService.AddTask(task)

}

func (controller *TaskController) RemoveTask(taskID string) error {
	return controller.taskService.RemoveTask(taskID)
}

func (controller *TaskController) GetTaskById(taskID string) (models.Task, error) {
	return controller.taskService.GetTaskById(taskID)
}

func (controller *TaskController) GetAllTasks() ([]models.Task, error) {
	return controller.taskService.GetAllTasks()
}

func (controller *TaskController) UpdateTask(taskID string, task models.Task) error {
	return controller.taskService.UpdateTask(taskID, task)
}

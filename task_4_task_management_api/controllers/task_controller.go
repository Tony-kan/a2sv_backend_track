package controllers

import (
	"errors"
	"net/http"
	services "task_4_task_management_api/data"
	"task_4_task_management_api/models"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService services.TaskServices
}

func NewTaskController(service services.TaskServices) *TaskController {
	return &TaskController{
		taskService: service,
	}
}

func (controller *TaskController) AddTask(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// createdTask, err := controller.taskService.AddTask(task)
	err := controller.taskService.AddTask(task)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task added successfully"})

	// ctx.JSON(http.StatusCreated, gin.H{"message": "Task added successfully", "task_id": createdTask.ID})

}

func (controller *TaskController) RemoveTask(ctx *gin.Context) {
	taskID := ctx.Param("id")
	err := controller.taskService.RemoveTask(taskID)
	if err != nil {
		if errors.Is(err, services.ErrTaskNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task removed successfully"})
}

func (controller *TaskController) GetTaskById(ctx *gin.Context) {
	taskID := ctx.Param("id")
	task, err := controller.taskService.GetTaskById(taskID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (controller *TaskController) GetAllTasks(ctx *gin.Context) {
	// return controller.taskService.GetAllTasks()
	tasks, err := controller.taskService.GetAllTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"number_of_tasks": len(tasks), "tasks": tasks})
}

func (controller *TaskController) UpdateTask(ctx *gin.Context) {
	// return controller.taskService.UpdateTask(taskID, task)
	taskID := ctx.Param("id")
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := controller.taskService.UpdateTask(taskID, task); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

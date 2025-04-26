package controllers

import (
	"errors"
	"net/http"
	services "task_6_task_management_api_with_auth/data"
	"task_6_task_management_api_with_auth/models"
	"time"

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

// AddTask godoc
// @Summary Create a new task
// @Description Add a new task to the system
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param   task body models.Task true "Task details"
// @Success 201 {object} map[string]interface{} "Task created"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/tasks [post]
func (controller *TaskController) AddTask(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// createdTask, err := controller.taskService.AddTask(task)
	taskID, err := controller.taskService.AddTask(task)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task added successfully", "task_id": taskID})

}

// RemoveTask godoc
// @Summary Delete a task by ID
// @Description Delete a task by its ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param   id path string true "Task ID"
// @Success 200 {object} map[string]interface{} "Task deleted"
// @Failure 404 {object} map[string]interface{} "Task not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/tasks/{id} [delete]
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

// GetTaskById godoc
// @Summary Get a task by ID
// @Description Get a task by its ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param   id path string true "Task ID"
// @Success 200 {object} models.Task "Task details"
// @Failure 404 {object} map[string]interface{} "Task not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/tasks/{id} [get]
func (controller *TaskController) GetTaskById(ctx *gin.Context) {
	taskID := ctx.Param("id")
	task, err := controller.taskService.GetTaskById(taskID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

// GetAllTasks godoc
// @Summary Get all tasks
// @Description Get a list of all tasks
// @Tags tasks
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{} "List of tasks"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/tasks/ [get]
func (controller *TaskController) GetAllTasks(ctx *gin.Context) {
	// return controller.taskService.GetAllTasks()
	tasks, err := controller.taskService.GetAllTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"number_of_tasks": len(tasks), "tasks": tasks})
}

// UpdateTask godoc
// @Summary Update a task by ID
// @Description Update a task by its ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param   id path string true "Task ID"
// @Param   task body models.Task true "Updated task details"
// @Success 200 {object} map[string]interface{} "Task updated"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 404 {object} map[string]interface{} "Task not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/tasks/{id} [put]
func (controller *TaskController) UpdateTask(ctx *gin.Context) {
	// return controller.taskService.UpdateTask(taskID, task)
	taskID := ctx.Param("id")
	//var task models.Task
	var updateData map[string]interface{}

	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if dueDateStr, exists := updateData["due_date"]; exists {
		if dateStr, ok := dueDateStr.(string); ok {
			parsedDate, err := time.Parse("2006-01-02", dateStr)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid due_date format, use YYYY-MM-DD"})
				return
			}
			updateData["due_date"] = parsedDate
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "due_date must be a string"})
			return
		}
	}

	validUpdate := make(map[string]interface{})

	for _, field := range []string{"title", "description", "status", "due_date"} {
		if value, exists := updateData[field]; exists {
			validUpdate[field] = value
		}
	}

	if len(validUpdate) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no valid fields provided for update"})
		return
	}
	if err := controller.taskService.UpdateTask(taskID, validUpdate); err != nil {
		if errors.Is(err, services.ErrTaskNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

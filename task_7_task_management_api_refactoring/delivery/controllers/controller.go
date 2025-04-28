package controllers

import (
	"errors"
	"net/http"
	domain "task_7_task_management_api_refactoring/domain"

	// services "task_7_task_management_api_refactoring/data"

	// "task_7_task_management_api_refactoring/models"
	"time"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	// taskService services.TaskServices
	// userService services.UserServices
	taskUsecase domain.TaskUsecase
	userUsecase domain.UserUsecase
}

// func NewController(taskService services.TaskServices, userService services.UserServices) *Controller {
// 	return &Controller{
// 		taskService: taskService,
// 		userService: userService,
// 	}
// }

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
func (controller *Controller) AddTask(ctx *gin.Context) {
	var task domain.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// createdTask, err := controller.taskService.AddTask(task)
	// taskID, err := controller.taskService.AddTask(task)
	taskID, err := controller.taskUsecase.AddTask(ctx, &task)

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
func (controller *Controller) RemoveTask(ctx *gin.Context) {
	taskID := ctx.Param("id")
	// err := controller.taskService.RemoveTask(taskID)
	err := controller.taskUsecase.RemoveTask(ctx, taskID)

	if err != nil {
		if errors.Is(err, domain.ErrTaskNotFound) {
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
func (controller *Controller) GetTaskById(ctx *gin.Context) {
	taskID := ctx.Param("id")
	// task, err := controller.taskService.GetTaskById(taskID)
	task, err := controller.taskUsecase.GetTaskById(ctx, taskID)

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
func (controller *Controller) GetAllTasks(ctx *gin.Context) {
	// tasks, err := controller.taskService.GetAllTasks()
	tasks, err := controller.taskUsecase.GetAllTasks(ctx)

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
func (controller *Controller) UpdateTask(ctx *gin.Context) {
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
	// if err := controller.taskService.UpdateTask(taskID, validUpdate); err != nil {
	if err := controller.taskUsecase.UpdateTask(ctx, taskID, validUpdate); err != nil {

		if errors.Is(err, domain.ErrTaskNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

// RegisterUser godoc
// @Summary Create a new user
// @Description Add a new user to the system
// @Tags users
// @Accept  json
// @Produce  json
// @Param   task body models.User true "User details"
// @Success 201 {object} map[string]interface{} "User register"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/users [post]
func (controller *Controller) RegisterUser(ctx *gin.Context) {
	// var task models.Task
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// createdTask, err := controller.taskService.AddTask(task)
	// taskID, err := controller.taskService.AddTask(task)
	// userId, err := controller.userService.RegisterUser(user)
	userId, err := controller.userUsecase.RegisterUser(ctx, &user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User Registered successfully", "user_id": userId})

}

// LoginUser godoc
// @Summary Login a user
// @Description login a user to the system
// @Tags users
// @Accept  json
// @Produce  json
// @Param   task body models.User true "User details"
// @Success 201 {object} map[string]interface{} "User logged in "
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/users [post]
func (controller *Controller) LoginUser(ctx *gin.Context) {
	// var task models.Task
	// var user models.User
	var loginRequest domain.LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// createdTask, err := controller.taskService.AddTask(task)
	// taskID, err := controller.taskService.AddTask(task)
	// userId, err := controller.userService.RegisterUser(user)
	// token, err := controller.userService.LoginUser(loginRequest)
	token, err := controller.userUsecase.LoginUser(ctx, loginRequest)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User Logged in successfully", "token": token})

}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{} "List of users"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/tasks/ [get]
func (controller *Controller) GetAllUsers(ctx *gin.Context) {

	// users, err := controller.userService.GetAllUsers()
	users, err := controller.userUsecase.GetAllUsers(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"number_of_users": len(users), "users": users})
}

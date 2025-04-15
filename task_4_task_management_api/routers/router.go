package routers

import (
	"net/http"
	"task_4_task_management_api/controllers"
	"task_4_task_management_api/models"

	"github.com/gin-gonic/gin"
)

type TaskRouter struct {
	taskController *controllers.TaskController
}

func NewTaskRouter(taskController *controllers.TaskController) *TaskRouter {
	return &TaskRouter{
		taskController: taskController,
	}
}

func (taskRouter *TaskRouter) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1/tasks")
	{
		api.POST("", taskRouter.AddTask)
		api.DELETE("/:id", taskRouter.RemoveTask)
		api.GET("/:id", taskRouter.GetTaskById)
		api.GET("", taskRouter.GetAllTasks)
		api.PUT("/:id", taskRouter.UpdateTask)
		api.PATCH("/:id", taskRouter.UpdateTask)
	}
}

func (taskRouter *TaskRouter) AddTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := taskRouter.taskController.AddTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Task added successfully", "task_id": task.ID})
}

func (taskRouter *TaskRouter) RemoveTask(c *gin.Context) {
	taskID := c.Param("id")
	err := taskRouter.taskController.RemoveTask(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task removed successfully"})
}

func (taskRouter *TaskRouter) GetTaskById(c *gin.Context) {
	taskID := c.Param("id")
	task, err := taskRouter.taskController.GetTaskById(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (taskRouter *TaskRouter) GetAllTasks(c *gin.Context) {
	tasks, err := taskRouter.taskController.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"number of tasks": len(tasks), "tasks": tasks})
}

func (taskRouter *TaskRouter) UpdateTask(c *gin.Context) {
	taskID := c.Param("id")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := taskRouter.taskController.UpdateTask(taskID, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

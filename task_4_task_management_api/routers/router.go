package routers

import (
	// "net/http"
	"task_4_task_management_api/controllers"
	// "task_4_task_management_api/models"

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
		api.POST("", taskRouter.taskController.AddTask)
		api.DELETE("/:id", taskRouter.taskController.RemoveTask)
		api.GET("/:id", taskRouter.taskController.GetTaskById)
		api.GET("", taskRouter.taskController.GetAllTasks)
		api.PUT("/:id", taskRouter.taskController.UpdateTask)
		api.PATCH("/:id", taskRouter.taskController.UpdateTask)
	}
}

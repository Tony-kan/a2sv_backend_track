package routers

import (
	"task_8_task_management_api_testing/delivery/controllers"
	Infrastructure "task_8_task_management_api_testing/infrastructure"

	"github.com/gin-gonic/gin"
)

type Router struct {
	generalController *controllers.Controller
}

func NewRouter(controller *controllers.Controller) *Router {
	return &Router{
		generalController: controller,
	}
}

// Todo : only admin can access all users details
// Todo : all users can view,create,update & delete tasks

func (generalRouter *Router) SetupRoutes(router *gin.Engine) {
	taskEndpoints := router.Group("/api/v1/tasks")
	taskEndpoints.Use(Infrastructure.AuthMiddleware())
	{
		taskEndpoints.POST("", generalRouter.generalController.AddTask)
		taskEndpoints.DELETE("/:id", generalRouter.generalController.RemoveTask)
		taskEndpoints.GET("/:id", generalRouter.generalController.GetTaskById)
		taskEndpoints.GET("", generalRouter.generalController.GetAllTasks)
		taskEndpoints.PUT("/:id", generalRouter.generalController.UpdateTask)
		taskEndpoints.PATCH("/:id", generalRouter.generalController.UpdateTask)
	}
	userEndpoints := router.Group("/api/v1/users")
	{
		userEndpoints.GET("", Infrastructure.AuthMiddleware(), Infrastructure.RequireRole("admin"), generalRouter.generalController.GetAllUsers)
		userEndpoints.POST("/register", generalRouter.generalController.RegisterUser)
		userEndpoints.POST("/login", generalRouter.generalController.LoginUser)
	}
}

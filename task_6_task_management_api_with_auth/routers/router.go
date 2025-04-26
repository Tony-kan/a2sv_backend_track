package routers

import (
	"task_6_task_management_api_with_auth/controllers"

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

func (generalRouter *Router) SetupRoutes(router *gin.Engine) {
	taskEndpoints := router.Group("/api/v1/tasks")
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
		userEndpoints.GET("", generalRouter.generalController.GetAllUsers)
		userEndpoints.POST("/register", generalRouter.generalController.RegisterUser)
		userEndpoints.POST("/login", generalRouter.generalController.LoginUser)
	}
}

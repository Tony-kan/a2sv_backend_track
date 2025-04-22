package main

import (
	"task_5_task_management_api_with_mongodb/controllers"
	services "task_5_task_management_api_with_mongodb/data"
	"task_5_task_management_api_with_mongodb/routers"

	"github.com/gin-gonic/gin"
	_ "task_5_task_management_api_with_mongodb/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	taskService := services.NewTaskService() // Ensure NewTaskService returns services.TaskService
	taskController := controllers.NewTaskController(taskService)
	taskRouter := routers.NewTaskRouter(taskController)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	taskRouter.SetupRoutes(router)

	router.Run("localhost:8080")
	// router.Run(":8080") // Start the server on port 8080
}

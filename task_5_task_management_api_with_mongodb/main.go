package main

import (
	"task_5_task_management_api_with_mongodb/controllers"
	services "task_5_task_management_api_with_mongodb/data"
	"task_5_task_management_api_with_mongodb/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	taskService := services.NewTaskService() // Ensure NewTaskService returns services.TaskService
	taskController := controllers.NewTaskController(taskService)
	taskRouter := routers.NewTaskRouter(taskController)

	router := gin.Default()
	taskRouter.SetupRoutes(router)

	router.Run("localhost:8080")
	// router.Run(":8080") // Start the server on port 8080
}

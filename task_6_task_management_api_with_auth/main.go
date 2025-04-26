package main

import (
	"context"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"task_6_task_management_api_with_auth/controllers"
	services "task_6_task_management_api_with_auth/data"
	"task_6_task_management_api_with_auth/routers"

	"github.com/gin-gonic/gin"
	_ "task_6_task_management_api_with_auth/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	//  MongoDB Connection
	godotenv.Load()
	mongo_uri := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongo_uri))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("taskdb")
	tasksCollection := db.Collection("tasks")

	taskService := services.NewTaskService(tasksCollection)

	taskController := controllers.NewTaskController(taskService)
	taskRouter := routers.NewTaskRouter(taskController)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	taskRouter.SetupRoutes(router)

	router.Run("localhost:8080")
	// router.Run(":8080") // Start the server on port 8080
}

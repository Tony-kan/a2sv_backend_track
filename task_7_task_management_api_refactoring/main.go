package main

import (
	"context"

	"github.com/joho/godotenv"

	"log"
	"os"
	"task_7_task_management_api_refactoring/controllers"
	services "task_7_task_management_api_refactoring/data"
	"task_7_task_management_api_refactoring/routers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "task_7_task_management_api_refactoring/docs"

	"github.com/gin-gonic/gin"

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
	userCollection := db.Collection("users")

	taskService := services.NewTaskService(tasksCollection)
	userService := services.NewUserService(userCollection)

	controller := controllers.NewController(taskService, userService)
	taskRouter := routers.NewRouter(controller)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	taskRouter.SetupRoutes(router)

	router.Run("localhost:8080")
	// router.Run(":8080") // Start the server on port 8080
}

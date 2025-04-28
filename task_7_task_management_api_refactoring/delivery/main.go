package main

import (
	"context"
	"time"

	"github.com/joho/godotenv"

	"log"
	"os"
	"task_7_task_management_api_refactoring/Delivery/routers"
	domain "task_7_task_management_api_refactoring/Domain"
	"task_7_task_management_api_refactoring/delivery/controllers"
	"task_7_task_management_api_refactoring/repositories"
	"task_7_task_management_api_refactoring/usecases"

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
	// tasksCollection := db.Collection("tasks")
	// userCollection := db.Collection("users")
	taskRepository := repositories.NewTaskRepository(db, domain.CollectionTask)
	taskUsecase := usecases.NewTaskUsecase(taskRepository, 10*time.Second)

	userRepository := repositories.NewUserRepository(db, domain.CollectionUser)
	userUsecase := usecases.NewUserUsecase(userRepository, 10*time.Second)

	// taskService := services.NewTaskService(tasksCollection)
	// userService := services.NewUserService(userCollection)
	// taskRepository := repositories.NewTaskRepository(*db, tasksCollection)

	// controller := controllers.NewController(taskService, userService)
	controller := &controllers.Controller{
		taskUsecase: taskUsecase,
		userUsecase: userUsecase,
	}
	taskRouter := routers.NewRouter(controller)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	taskRouter.SetupRoutes(router)

	router.Run("localhost:8080")
	// router.Run(":8080") // Start the server on port 8080
}

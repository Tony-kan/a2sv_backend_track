package main

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"

	"log"
	"task_8_task_management_api_testing/delivery/controllers"
	"task_8_task_management_api_testing/delivery/routers"

	domain "task_8_task_management_api_testing/domain"
	"task_8_task_management_api_testing/repositories"
	"task_8_task_management_api_testing/usecases"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "task_8_task_management_api_testing/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	//  MongoDB Connection
	// godotenv.Load("../.env")
	godotenv.Load("./.env")
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

	taskRepository := repositories.NewTaskRepository(db, domain.CollectionTask)

	userRepository := repositories.NewUserRepository(db, domain.CollectionUser)

	userUC := usecases.NewUserUsecase(userRepository, 10*time.Second)
	taskUC := usecases.NewTaskUsecase(taskRepository, 10*time.Second)

	controller := &controllers.Controller{
		TaskUsecase: taskUC,
		UserUsecase: userUC,
	}
	taskRouter := routers.NewRouter(controller)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	taskRouter.SetupRoutes(router)

	router.Run("localhost:8080")
	// router.Run(":8080") // Start the server on port 8080
}

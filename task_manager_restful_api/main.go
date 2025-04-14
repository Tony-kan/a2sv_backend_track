package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @title Task Manager API
// @version 1.0
// @description This is a simple task manager API.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support

// Task represent a task with its properties
// Todo: create endpoints to get all tasks, get a task by id,
// Todo: create endpoints to create a new task, and update a task by id

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

// Mock data for tasks
var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "Description for Task 1", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Description for Task 2", DueDate: time.Now().AddDate(0, 0, 1), Status: "Completed"},
	{ID: "3", Title: "Task 3", Description: "Description for Task 3", DueDate: time.Now().AddDate(0, 0, 2), Status: "In Progress"},
}

func main() {
	router := gin.Default()

	//  a router to ping
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// a router to get all tasks
	router.GET("/tasks", getAllTasks)

	// a router to get a task by id
	router.GET("/tasks/:id", getTaskById)

	// a router to create a new task
	router.POST("/tasks", createTask)

	// a router to update a task by id
	router.PUT("/tasks/:id", updateTaskById)

	// a router to delete a task by id
	router.DELETE("/tasks/:id", deleteTaskById)

	router.Run("localhost:8080")
}

func getAllTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

func getTaskById(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, task := range tasks {
		if task.ID == id {
			ctx.JSON(http.StatusOK, task)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Task not found"})
}

func createTask(ctx *gin.Context) {
	var newTask Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	tasks = append(tasks, newTask)
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
		"task":    newTask,
	})
}

func updateTaskById(ctx *gin.Context) {
	id := ctx.Param("id")

	var updatedTask Task

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			// Update the task on the specified fields
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			if updatedTask.Status != "" {
				tasks[i].Status = updatedTask.Status
			}
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Task updated successfully",
				"task":    tasks[i],
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Task not found"})
}
func deleteTaskById(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, val := range tasks {
		if val.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Task deleted successfully",
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Task not found"})
}

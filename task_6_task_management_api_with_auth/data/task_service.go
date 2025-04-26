package services

import (
	"context"
	"errors"
	"fmt"
	"task_6_task_management_api_with_auth/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
)

// Todo : create an interface,services and constructor for the task service
// Todo : implement the methods of the interface
// Todo : create error handling methods
// Todo : Replace in-memory storage

var (
	ErrTaskNotFound  = errors.New("task not found")
	ErrTaskExists    = errors.New("task already exists")
	ErrInvalidTask   = errors.New("invalid task data")
	ErrInvalidTaskID = errors.New("invalid task ID")
)

type TaskServices interface {
	AddTask(task models.Task) (string, error)
	RemoveTask(taskID string) error
	GetTaskById(taskID string) (models.Task, error)
	GetAllTasks() ([]models.Task, error)
	UpdateTask(taskID string, updateFields map[string]interface{}) error
}

type TaskService struct {
	tasksCollection *mongo.Collection
}

func NewTaskService(tasksCollection *mongo.Collection) TaskServices {
	// Create index on createdAt field
	ctx := context.Background()
	indexModel := mongo.IndexModel{
		Keys: bson.D{{"createdAt", 1}},
	}
	if _, err := tasksCollection.Indexes().CreateOne(ctx, indexModel); err != nil {
		panic(fmt.Sprintf("Failed to create index: %v", err))
	}

	return &TaskService{
		tasksCollection: tasksCollection,
	}
}

func (service *TaskService) AddTask(task models.Task) (string, error) {

	if task.Title == "" {
		return "", ErrInvalidTask
	}
	task.ID = primitive.NewObjectID()
	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	if task.DueDate.IsZero() {
		task.DueDate = now.AddDate(0, 0, 7)
	}

	if task.Status == "" {
		task.Status = "Pending"
	}

	_, err := service.tasksCollection.InsertOne(context.Background(), task)
	if err != nil {
		return "", fmt.Errorf("failed to create task: %v", err)
	}

	return task.ID.Hex(), nil
}

func (service *TaskService) RemoveTask(taskID string) error {
	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return ErrInvalidTaskID
	}

	res, err := service.tasksCollection.DeleteOne(context.Background(), bson.M{"_id": objectID})

	if err != nil {
		return fmt.Errorf("failed to remove task: %v", err)
	}

	if res.DeletedCount == 0 {
		return ErrTaskNotFound
	}

	return nil
}

func (service *TaskService) GetTaskById(taskID string) (models.Task, error) {
	var task models.Task
	objID, err := primitive.ObjectIDFromHex(taskID)

	if err != nil {
		return task, ErrInvalidTaskID
	}

	err = service.tasksCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&task)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return task, ErrTaskNotFound
		}
		return task, fmt.Errorf("failed to get task: %v", err)
	}

	return task, err
}

func (service *TaskService) GetAllTasks() ([]models.Task, error) {
	//var tasks []models.Task
	tasks := make([]models.Task, 0)

	cur, err := service.tasksCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get all tasks: %v", err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var task models.Task
		if err := cur.Decode(&task); err != nil {
			return nil, fmt.Errorf("failed to decode task: %v", err)
		}
		tasks = append(tasks, task)
	}
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("failed to get all tasks: %v", err)
	}

	return tasks, nil
}

func (service *TaskService) UpdateTask(taskID string, updateFields map[string]interface{}) error {
	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return ErrInvalidTaskID
	}

	updateFields["updated_at"] = time.Now()

	res, err := service.tasksCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": objectID},
		bson.M{"$set": updateFields},
	)
	if err != nil {
		return fmt.Errorf("failed to update task: %v", err)
	}
	if res.MatchedCount == 0 {
		return ErrTaskNotFound
	}

	return nil
}

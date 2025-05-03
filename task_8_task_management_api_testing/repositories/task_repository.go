package repositories

import (
	"context"
	"errors"
	"fmt"
	"task_8_task_management_api_testing/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskRepository struct {
	database   *mongo.Database
	collection string
}

func NewTaskRepository(database *mongo.Database, collection string) domain.TaskRepository {
	return &taskRepository{
		database:   database,
		collection: collection,
	}
}

func (repository *taskRepository) GetCollection() *mongo.Collection {
	return repository.database.Collection(repository.collection)
}

func (repository *taskRepository) AddTask(ctx context.Context, task *domain.Task) (string, error) {

	if task.Title == "" {
		return "", domain.ErrInvalidTask
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

	_, err := repository.GetCollection().InsertOne(ctx, task)
	if err != nil {
		return "", fmt.Errorf("failed to create task: %v", err)
	}

	return task.ID.Hex(), nil
}

func (repository *taskRepository) RemoveTask(ctx context.Context, taskID string) error {
	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return domain.ErrInvalidTaskID
	}

	res, err := repository.GetCollection().DeleteOne(ctx, bson.M{"_id": objectID})

	if err != nil {
		return fmt.Errorf("failed to remove task: %v", err)
	}

	if res.DeletedCount == 0 {
		return domain.ErrTaskNotFound
	}

	return nil
}

func (repository *taskRepository) GetTaskById(ctx context.Context, taskID string) (*domain.Task, error) {
	var task domain.Task
	objID, err := primitive.ObjectIDFromHex(taskID)

	if err != nil {
		return &task, domain.ErrInvalidTaskID
	}

	err = repository.GetCollection().FindOne(ctx, bson.M{"_id": objID}).Decode(&task)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &task, domain.ErrTaskNotFound
		}
		return &task, fmt.Errorf("failed to get task: %v", err)
	}

	return &task, err
}

func (repository *taskRepository) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	//var tasks []models.Task
	tasks := make([]*domain.Task, 0)

	cur, err := repository.GetCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get all tasks: %v", err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var task domain.Task
		if err := cur.Decode(&task); err != nil {
			return nil, fmt.Errorf("failed to decode task: %v", err)
		}
		tasks = append(tasks, &task)
	}
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("failed to get all tasks: %v", err)
	}

	return tasks, nil
}

func (repository *taskRepository) UpdateTask(ctx context.Context, taskID string, updateFields map[string]interface{}) error {
	objectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return domain.ErrInvalidTaskID
	}

	updateFields["updated_at"] = time.Now()

	res, err := repository.GetCollection().UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": updateFields},
	)
	if err != nil {
		return fmt.Errorf("failed to update task: %v", err)
	}
	if res.MatchedCount == 0 {
		return domain.ErrTaskNotFound
	}

	return nil
}

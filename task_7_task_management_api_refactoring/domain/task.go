package domain

import (
	"context"
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CollectionTask = "tasks"

type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Status      string             `json:"status" bson:"status"`
	DueDate     time.Time          `json:"due_date" bson:"due_date"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type TaskRepository interface {
	AddTask(ctx context.Context, task *Task) (string, error)
	GetAllTasks(ctx context.Context) ([]*Task, error)
	GetTaskById(ctx context.Context, taskID string) (*Task, error)
	RemoveTask(ctx context.Context, taskID string) error
	UpdateTask(ctx context.Context, taskID string, updateFields map[string]interface{}) error
}

type TaskUsecase interface {
	AddTask(ctx context.Context, task *Task) (string, error)
	GetAllTasks(ctx context.Context) ([]*Task, error)
	GetTaskById(ctx context.Context, taskID string) (*Task, error)
	RemoveTask(ctx context.Context, taskID string) error
	UpdateTask(ctx context.Context, taskID string, updateFields map[string]interface{}) error
}

// UnmarshalJSON is a custom JSON unmarshaller for the Task struct.
// It handles the parsing of the DueDate field from a string to a time.Time object.
// The expected format for the date is "2006-01-02".
// If the date is not in the expected format, it returns an error.
func (t *Task) UnmarshalJSON(data []byte) error {
	type Alias Task
	aux := &struct {
		DueDate string `json:"due_date"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parse the date if provided
	if aux.DueDate != "" {
		parsedDate, err := time.Parse("2006-01-02", aux.DueDate)
		if err != nil {
			return err
		}
		t.DueDate = parsedDate
	}

	return nil
}

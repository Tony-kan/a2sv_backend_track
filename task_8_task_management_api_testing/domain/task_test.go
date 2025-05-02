package domain_test

import (
	"encoding/json"
	"task_8_task_management_api_testing/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTaskValidation(t *testing.T) {
	t.Run("ValidTask", func(t *testing.T) {

		input := `{
			"title": "Valid Task",
			"description": "Test Description",
			"status": "pending",
			"due_date": "2024-01-01"
		}`

		var task domain.Task
		err := json.Unmarshal([]byte(input), &task)

		assert.NoError(t, err)
		assert.Equal(t, "Valid Task", task.Title)
		assert.Equal(t, "Test Description", task.Description)
		assert.Equal(t, "pending", task.Status)
		assert.Equal(t, time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), task.DueDate)
	})
	t.Run("InvalidDateJSON", func(t *testing.T) {
		input := `{
			"title": "Invalid Task",
			"due_date": "2024-13-01"
		}`

		var task domain.Task
		err := json.Unmarshal([]byte(input), &task)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "parsing time")
	})

	t.Run("EmptyTitle", func(t *testing.T) {
		task := &domain.Task{
			ID:        primitive.NewObjectID(),
			Title:     "", // Empty title
			Status:    "pending",
			DueDate:   time.Now().Add(24 * time.Hour),
			CreatedAt: time.Now(),
		}

		assert.Empty(t, task.Title)
	})
}

package task

import (
	"testing"
	"time"
)

// Start with the domain model, smallest unit of value
// "I can create a new task, it starts incomplete and has a description"

func TestNewTask(t *testing.T) {
	t.Run("Creates a new task", func(t *testing.T) {
		task := NewTask("buy milk", 1, time.Date(2025, 8, 31, 12, 0, 0, 0, time.UTC))
		wantDescription := "buy milk"

		if task.Description != wantDescription {
			t.Errorf("Description: got %q, want %q", task.Description, wantDescription)
		}

		if task.Completed {
			t.Error("New task should not be completed")
		}
	})
}

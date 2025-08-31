package task

import (
	"testing"
	"time"
)

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

func TestAddToListOfTasks(t *testing.T) {
	t.Run("auto increments ID", func(t *testing.T) {
		// User should only configure the min amount they need to i.e. just description
		// all other stateful variables should be handled by business logic: ID, CreatedAt, Completed (when adding a task)
		list := TaskList{}
		list.AddTask("buy milk")
		list.AddTask("buy bread")

		if len(list.Tasks) != 2 {
			t.Error("failed to add 2 tasks")
		}
		if list.Tasks[0].ID != 1 {
			t.Errorf("got ID: %d, want ID: %d from task %q", list.Tasks[0].ID, 1, list.Tasks[0].Description)
		}
		if list.Tasks[1].ID != 2 {
			t.Errorf("got ID: %d, want ID: %d from task %q", list.Tasks[1].ID, 2, list.Tasks[1].Description)
		}
	})
}

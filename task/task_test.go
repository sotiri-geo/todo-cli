package task

import (
	"strings"
	"testing"
	"time"
)

func TestTask(t *testing.T) {
	t.Run("creates a new task", func(t *testing.T) {
		task := NewTask("buy milk", 1, time.Date(2025, 8, 31, 12, 0, 0, 0, time.UTC))
		wantDescription := "buy milk"

		if task.Description != wantDescription {
			t.Errorf("Description: got %q, want %q", task.Description, wantDescription)
		}

		if task.Completed {
			t.Error("New task should not be completed")
		}
	})

	t.Run("mark task as done", func(t *testing.T) {

		task := NewTask("buy milk", 1, time.Date(2025, 8, 31, 12, 0, 0, 0, time.UTC))

		task.MarkDone()

		if !task.Completed {
			t.Error("task failed to be marked as completed.")
		}
	})
}

func TestListOfTasks(t *testing.T) {
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

	t.Run("delete a task from list", func(t *testing.T) {
		list := TaskList{}
		task := list.AddTask("buy milk")

		err := list.DeleteTask(task.ID)

		if err != nil {
			t.Fatal("should not error.")
		}
		if len(list.Tasks) != 0 {
			t.Error("did not delete task")
		}
	})

	t.Run("task to delete not found", func(t *testing.T) {
		list := TaskList{}
		task := list.AddTask("buy milk")

		err := list.DeleteTask(task.ID + 1)

		if err == nil {
			t.Fatal("should error but didn't.")
		}

		if !strings.Contains(err.Error(), "not found") {
			t.Errorf("unexpected error type: %v", err)
		}
	})

	t.Run("tasks ordered by creation", func(t *testing.T) {
		list := TaskList{}
		task1 := list.AddTask("buy milk")
		task2 := list.AddTask("buy bread")

		list.DeleteTask(task2.ID)
		task3 := list.AddTask("buy cheese")

		if !task1.CreatedAt.Before(task3.CreatedAt) {
			t.Error("task 1 should be created before task 3")
		}
	})
}

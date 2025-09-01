package task

import (
	"errors"
	"strings"
	"testing"
	"time"
)

func TestTask(t *testing.T) {
	t.Run("creates a new task", func(t *testing.T) {
		task, err := NewTask("buy milk", 1, time.Date(2025, 8, 31, 12, 0, 0, 0, time.UTC))
		wantDescription := "buy milk"

		if err != nil {
			// NOTE TO SELF: use t.Fatal() when the failure makes rest of tests meaningless and need to stop program
			t.Fatal("should not fail")
		}

		if task.Description != wantDescription {
			t.Errorf("Description: got %q, want %q", task.Description, wantDescription)
		}

		if task.Completed {
			t.Error("New task should not be completed")
		}
	})

	t.Run("mark task as done", func(t *testing.T) {

		task, err := NewTask("buy milk", 1, time.Date(2025, 8, 31, 12, 0, 0, 0, time.UTC))

		if err != nil {
			t.Fatal("should not fail")
		}
		task.Complete()

		if !task.Completed {
			t.Error("task failed to be marked as completed.")
		}
	})

	t.Run("description should be non empty", func(t *testing.T) {

		_, err := NewTask("", 1, time.Date(2025, 8, 31, 12, 0, 0, 0, time.UTC))

		if err == nil {
			t.Fatal("should error")
		}

		if !errors.Is(err, ErrEmptyTaskDescription) {
			t.Errorf("got %q, want %q", err, ErrEmptyTaskDescription)
		}
	})
}

func TestListOfTasks(t *testing.T) {
	t.Run("auto increments ID", func(t *testing.T) {
		// User should only configure the min amount they need to i.e. just description
		// all other stateful variables should be handled by business logic: ID, CreatedAt, Completed (when adding a task)
		list := NewTaskList()
		_, err1 := list.AddTask("buy milk")
		_, err2 := list.AddTask("buy bread")

		if err1 != nil || err2 != nil {
			t.Fatal("should not error")
		}

		if len(list.Tasks) != 2 {
			t.Fatal("failed to add 2 tasks")
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
		task, _ := list.AddTask("buy milk")

		err := list.DeleteTask(task.ID)

		if err != nil {
			t.Fatal("should not error.")
		}
		if len(list.Tasks) != 0 {
			t.Error("did not delete task")
		}
	})

	t.Run("task to delete not found", func(t *testing.T) {
		list := NewTaskList()
		task, _ := list.AddTask("buy milk")

		err := list.DeleteTask(task.ID + 1)

		if err == nil {
			t.Fatal("should error but didn't.")
		}

		if !strings.Contains(err.Error(), "not found") {
			t.Errorf("unexpected error type: %v", err)
		}
	})

	t.Run("tasks ordered by creation", func(t *testing.T) {
		list := NewTaskList()
		task1, _ := list.AddTask("buy milk")
		time.Sleep(time.Millisecond) // ensure diff timestamps
		task2, _ := list.AddTask("buy bread")

		list.DeleteTask(task2.ID)
		time.Sleep(time.Millisecond)
		task3, _ := list.AddTask("buy cheese")

		if !task1.CreatedAt.Before(task3.CreatedAt) {
			t.Error("task 1 should be created before task 3")
		}
	})

	t.Run("unique ID creation for each task", func(t *testing.T) {
		list := NewTaskList()
		task1, _ := list.AddTask("buy milk")
		task2, _ := list.AddTask("buy bread")

		list.DeleteTask(task1.ID)

		// Create new ID should have id different from task2
		task3, _ := list.AddTask("buy cheese")

		if task2.ID == task3.ID {
			t.Errorf("previous task Id %d, new task Id %d ", task2.ID, task3.ID)
		}
	})

	t.Run("mark task as done after adding to list", func(t *testing.T) {
		list := NewTaskList()
		task, _ := list.AddTask("buy milk")

		task.Complete()

		// Get the task from list
		taskInList := list.Tasks[0]

		if !taskInList.Completed {
			t.Error("task was not marked as complete")
		}
	})

	t.Run("find all completed tasks", func(t *testing.T) {
		list := NewTaskList()

		list.AddTask("buy milk")
		want, _ := list.AddTask("buy bread") // represents task to be marked as completed

		want.Complete()

		if len(list.FindCompleted()) != 1 {
			t.Fatal("should have one completed tasks")
		}

		got := list.FindCompleted()[0]
		if got != want {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})
	t.Run("find task by id", func(t *testing.T) {
		list := NewTaskList()
		task, _ := list.AddTask("Buy milk")
		list.AddTask("Buy bread")

		got, err := list.GetTask(task.ID)

		if err != nil {
			t.Fatalf("Failed to get task: %v", err)
		}

		if got != task {
			t.Errorf("got %v, want %v", got, task)
		}
	})

	t.Run("search for unknown id", func(t *testing.T) {
		list := NewTaskList()

		_, err := list.GetTask(1)

		if err == nil {
			t.Fatal("should have errored")
		}

		if !errors.Is(err, ErrNotFoundTask) {
			t.Errorf("got %v, want %v", err, ErrNotFoundTask)
		}
	})

	t.Run("mark task as done by id", func(t *testing.T) {
		list := NewTaskList()
		task, _ := list.AddTask("Buy milk")
		list.AddTask("Buy bread")

		taskCompleted, err := list.MarkCompleted(task.ID)

		if err != nil {
			t.Fatalf("Unable to mark task as completed: Found %v", err)
		}

		if !taskCompleted.Completed {
			t.Fatal("Did not mark task as completed")
		}

		// Check integrity
		listCompleted := list.FindCompleted()
		if len(listCompleted) != 1 {
			t.Fatalf("Should have 1 completed tasks: Found %d", len(listCompleted))
		}

		if listCompleted[0] != task {
			t.Errorf("Incorrect task marked as complete: got %v, want %v", listCompleted[0], task)
		}
	})
}

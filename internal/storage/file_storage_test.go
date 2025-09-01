package storage

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/sotiri-geo/todo-cli/internal/task"
)

// We need to tests components that touch the file system

func TestFileStorage_Integration(t *testing.T) {
	// Create temporary directory for test files
	tempDir, err := os.MkdirTemp("", "todo-test-*")

	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	defer os.RemoveAll(tempDir)

	t.Run("save and load from actual file", func(t *testing.T) {
		filename := filepath.Join(tempDir, "tasks.json")
		store := NewFileStore(filename)

		// Create test data
		originalList := task.NewTaskList()
		task1, _ := originalList.AddTask("Buy milk")
		originalList.AddTask("Buy bread")

		task1.Complete()

		// Save to file
		err := store.Save(originalList)

		if err != nil {
			t.Fatalf("Failed to save: %v", err)
		}
		// Verify file was persisted
		if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
			t.Fatalf("File does not exist: %v", err)
		}

		// Reload and check same content as original
		loadedList, err := store.Load()

		if err != nil {
			t.Errorf("Load failed: %v", err)
		}

		// Check loaded content, individually

		if len(originalList.Tasks) != len(loadedList.Tasks) {
			t.Fatalf("Failed to load tasks: got %v, want %v", len(loadedList.Tasks), len(originalList.Tasks))
		}

		for i, originalTask := range originalList.Tasks {
			assertTasksEqual(t, originalTask, loadedList.Tasks[i])
		}

		if len(loadedList.FindCompleted()) != 1 {
			t.Errorf("Should have %d completed tasks", 1)
		}
	})

	t.Run("save, load, delete", func(t *testing.T) {

		filename := filepath.Join(tempDir, "tasks.json")
		store := NewFileStore(filename)

		// Create test data
		originalList := task.NewTaskList()
		task1, _ := originalList.AddTask("Buy milk")
		originalList.AddTask("Buy bread")

		// Save to file
		err := store.Save(originalList)

		if err != nil {
			t.Fatalf("Failed to save: %v", err)
		}
		// Verify file was persisted
		if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
			t.Fatalf("File does not exist: %v", err)
		}

		// Reload and check same content as original
		loadedList, err := store.Load()

		loadedList.DeleteTask(task1.ID)

		// Check IDs retain uniqueness
		loadedList.AddTask("Buy cheese")

		err = store.Save(loadedList)

		nextLoadedList, err := store.Load()

		assertUniqueIds(t, nextLoadedList)

	})
}

func assertTasksEqual(t testing.TB, got, want *task.Task) {
	t.Helper()

	if got.ID != want.ID {
		t.Errorf("ID mismatch: got %d, want %d", got.ID, want.ID)
	}

	if got.Description != want.Description {
		t.Errorf("Description mismatch: got %q, want %q", got.Description, want.Description)
	}

	if got.Completed != want.Completed {
		t.Errorf("Completed mismatch: got %t, want %t", got.Completed, want.Completed)
	}

	// handles monotonic clock differences
	if !got.CreatedAt.Equal(want.CreatedAt) {
		t.Errorf("CreatedAt mismatch: got %v, want %v", got.CreatedAt, want.CreatedAt)
	}
}

func assertUniqueIds(t testing.TB, taskList *task.TaskList) {

	t.Helper()
	seen := map[int]struct{}{}

	for _, task := range taskList.Tasks {
		if _, found := seen[task.ID]; !found {
			seen[task.ID] = struct{}{}
		} else {
			t.Fatalf("Found duplicate ID %d from description %q", task.ID, task.Description)
		}
	}
}

package task

import (
	"testing"

	"github.com/sotiri-geo/todo-cli/internal/task"
)

// Spystore needs to implement the interface Store
type SpyStore struct {
	taskList      task.TaskList
	SaveCallCount int
	LoadCallCount int
}

func (s *SpyStore) Save(taskList *task.TaskList) error {
	s.SaveCallCount++
	s.taskList = *taskList // store a copy of taskList to mimic persisting data
	return nil
}

func (s *SpyStore) Load() (*task.TaskList, error) {

	return &s.taskList, nil
}

func TestService(t *testing.T) {
	t.Run("add task", func(t *testing.T) {
		description := "Buy milk"
		store := &SpyStore{}
		svc := TaskService{store}
		got, err := svc.AddTask(description)
		want := task.NewTaskList()
		want.AddTask(description)

		if err != nil {
			t.Fatalf("should not error: found %v", err)
		}
		if store.SaveCallCount != 1 {
			t.Fatal("Did not call save once.")
		}

		// Check integrity
		assertEqualTaskLists(t, *got, *want)
	})
}

func assertEqualTaskLists(t testing.TB, got, want task.TaskList) {
	t.Helper()

	// Check integrity
	if len(got.Tasks) != len(want.Tasks) {
		t.Fatalf("Task list mismatch: got %d, want %d", len(got.Tasks), len(want.Tasks))
	}

	// Check data
	for i, gotTask := range got.Tasks {
		if gotTask.ID != want.Tasks[i].ID {
			t.Errorf("ID: got %d, want %d", gotTask.ID, want.Tasks[i].ID)
		}

		if gotTask.Description != want.Tasks[i].Description {
			t.Errorf("Description: got %q, want %q", gotTask.Description, want.Tasks[i].Description)
		}

		if gotTask.Completed != want.Tasks[i].Completed {
			t.Errorf("Completed: got %v, want %v", gotTask.Completed, want.Tasks[i].Completed)
		}
	}
}

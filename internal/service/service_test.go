package task

import (
	"testing"

	"github.com/sotiri-geo/todo-cli/internal/task"
)

// Spystore needs to implement the interface Store
type SpyStore struct {
	taskList               task.TaskList
	SaveCallCount          int
	LoadCallCount          int
	MarkCompletedCallCount int
}

func (s *SpyStore) Save(taskList *task.TaskList) error {
	s.SaveCallCount++
	s.taskList = *taskList // store a copy of taskList to mimic persisting data
	return nil
}

func (s *SpyStore) Load() (*task.TaskList, error) {
	s.LoadCallCount++
	return &s.taskList, nil
}

func TestService(t *testing.T) {
	t.Run("add task", func(t *testing.T) {
		description := "Buy milk"
		store := &SpyStore{}
		svc := TaskService{store}
		got, err := svc.AddTask(description)
		wantTaskList := task.NewTaskList()
		wantTask, _ := wantTaskList.AddTask(description)

		if err != nil {
			t.Fatalf("should not error: found %v", err)
		}
		if store.SaveCallCount != 1 {
			t.Fatal("Did not call save once.")
		}

		// Check integrity
		assertTasksEqual(t, got, wantTask)

	})

	t.Run("list all tasks", func(t *testing.T) {
		list := task.NewTaskList()
		list.AddTask("Buy milk")
		store := &SpyStore{taskList: *list}
		svc := TaskService{store}

		loaded, err := svc.ListTasks()

		if err != nil {
			t.Fatalf("Failed to load: %v", err)
		}

		if store.LoadCallCount != 1 {
			t.Fatal("Did not call load method once")
		}
		assertEqualTaskLists(t, *list, *loaded)

	})

	t.Run("list all completed tasks - optional flag", func(t *testing.T) {

		list := task.NewTaskList()
		task1, _ := list.AddTask("Buy milk")
		list.AddTask("Buy bread")

		task1.Complete()
		store := &SpyStore{taskList: *list}
		svc := TaskService{store}

		loaded, err := svc.ListCompletedTasks()

		if err != nil {
			t.Fatalf("Failed to load tasks: %v", err)
		}

		if len(loaded.Tasks) != 1 {
			t.Fatal("Failed to mark task as completed")
		}

		gotMarkedAsCompleted := loaded.Tasks[0]

		if gotMarkedAsCompleted != task1 {
			t.Errorf("Incorrect task marked as completed: got %+v, want %+v", gotMarkedAsCompleted, task1)
		}

	})

	t.Run("mark task as completed", func(t *testing.T) {
		list := task.NewTaskList()
		task1, _ := list.AddTask("Buy milk")
		list.AddTask("Buy bread")

		store := &SpyStore{taskList: *list}
		svc := TaskService{store}

		completedTask, err := svc.MarkCompleted(task1.ID)

		if err != nil {
			t.Fatalf("should not error: Found %v", err)
		}

		if task1 != completedTask {
			t.Errorf("got %+v, want %+v", completedTask, task1)
		}

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

func assertTasksEqual(t *testing.T, got, want *task.Task) {
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
}

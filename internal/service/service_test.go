package service

import (
	"errors"
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
		svc := NewTaskService(store)
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

	t.Run("add multiple tasks", func(t *testing.T) {
		store := &SpyStore{}
		svc := NewTaskService(store)
		task1, _ := svc.AddTask("Buy Milk")
		task2, _ := svc.AddTask("Buy Bread")

		// retains integrity

		if task1.ID == task2.ID {
			t.Errorf("Duplicate ID in both tasks: %d", task1.ID)
		}
	})

	t.Run("list all tasks", func(t *testing.T) {
		list := task.NewTaskList()
		list.AddTask("Buy milk")
		store := &SpyStore{taskList: *list}
		svc := NewTaskService(store)

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
		svc := NewTaskService(store)

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

	t.Run("list all pending tasks - optional flag", func(t *testing.T) {

		list := task.NewTaskList()
		task1, _ := list.AddTask("Buy milk")
		task2, _ := list.AddTask("Buy bread")

		task1.Complete()
		store := &SpyStore{taskList: *list}
		svc := NewTaskService(store)

		loaded, err := svc.ListPendingTasks()

		if err != nil {
			t.Fatalf("Failed to load tasks: %v", err)
		}
		pendingTasks := loaded.FindPending()
		if len(pendingTasks) != 1 {
			t.Fatal("Failed to mark task as completed")
		}

		if pendingTasks[0] != task2 {
			t.Errorf("Task should be pending: got %v, want %v", pendingTasks[0], task2)
		}
	})

	t.Run("mark task as completed", func(t *testing.T) {
		list := task.NewTaskList()
		task1, _ := list.AddTask("Buy milk")
		list.AddTask("Buy bread")

		store := &SpyStore{taskList: *list}
		svc := NewTaskService(store)

		completedTask, err := svc.MarkCompleted(task1.ID)

		if err != nil {
			t.Fatalf("should not error: Found %v", err)
		}

		if store.SaveCallCount != 1 {
			t.Fatalf("Failed to call save: Found count %d", store.SaveCallCount)
		}

		if task1 != completedTask {
			t.Errorf("got %+v, want %+v", completedTask, task1)
		}

		got, _ := list.GetTask(task1.ID)

		if !got.Completed {
			t.Errorf("task was not marked as completed: got %v", got)
		}
	})

	t.Run("get task by id", func(t *testing.T) {
		list := task.NewTaskList()
		task1, _ := list.AddTask("Buy milk")
		list.AddTask("Buy bread")

		store := &SpyStore{taskList: *list}
		svc := NewTaskService(store)

		gotTask, err := svc.GetTask(task1.ID)

		if err != nil {
			t.Fatal("should not error")
		}

		if gotTask != task1 {
			t.Errorf("got %v, want %v", gotTask, task1)
		}

	})

	t.Run("delete task", func(t *testing.T) {
		list := task.NewTaskList()
		task1, _ := list.AddTask("Buy milk")
		list.AddTask("Buy bread")

		store := &SpyStore{taskList: *list}
		svc := NewTaskService(store)

		err := svc.DeleteTask(task1.ID)

		if err != nil {
			t.Fatal("should not error")
		}

		loaded, err := svc.ListTasks()
		if err != nil {
			t.Fatal("should not error")
		}
		if len(loaded.Tasks) != 1 {
			t.Fatalf("Number of tasks: got %d, want %d", len(loaded.Tasks), 1)
		}

		// Make sure ID is no longer available
		_, errGet := svc.GetTask(task1.ID)

		if !errors.Is(errGet, task.ErrNotFoundTask) {
			t.Errorf("got %q, want %q", errGet, task.ErrNotFoundTask)
		}
	})
}

func TestService_Integration(t *testing.T) {

	t.Run("adding multiple tasks persists", func(t *testing.T) {
		store := &SpyStore{}
		svc := NewTaskService(store)

		task1, _ := svc.AddTask("Buy milk")
		task2, _ := svc.AddTask("Buy bread")

		want := task.TaskList{Tasks: []*task.Task{task1, task2}}
		wantCount := len(want.Tasks)

		loaded, _ := svc.ListTasks()

		if len(loaded.Tasks) != wantCount {
			t.Fatalf("Failed to persist tasks: got %d, want %d", len(loaded.Tasks), wantCount)
		}

		if store.SaveCallCount != wantCount {
			t.Fatalf("Failed to call save to store the required number of times: got %d, want %d", store.SaveCallCount, wantCount)
		}
		// Check integrity
		if loaded.Tasks[0].ID == loaded.Tasks[1].ID {
			t.Fatalf("Duplicate ID detected: %d", loaded.Tasks[0].ID)
		}

		assertEqualTaskLists(t, *loaded, want)

	})

	t.Run("multiple state operations", func(t *testing.T) {

		store := &SpyStore{}
		svc := NewTaskService(store)

		task, _ := svc.AddTask("Buy milk")
		svc.AddTask("Buy bread")
		svc.AddTask("Buy cheese")

		svc.DeleteTask(task.ID) // Remove first task
		loaded, _ := svc.ListTasks()

		if len(loaded.Tasks) != 2 {
			t.Errorf("Failed to persist tasks: got %d, want %d", len(loaded.Tasks), 3)
		}

		// Check integrity
		svc.AddTask("Buy water")

		loadedAgain, _ := svc.ListTasks()

		assertUniqueIds(t, loadedAgain)

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

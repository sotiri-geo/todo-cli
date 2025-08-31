package storage

import (
	"fmt"
	"testing"

	"github.com/sotiri-geo/todo-cli/internal/task"
)

// Fake, spy store for tracking behaviour
type FakeStore struct {
	Filename  string
	taskList  *task.TaskList
	LoadCalls int
	SaveCalls int
}

func (f *FakeStore) Load() (*task.TaskList, error) {
	f.LoadCalls++
	return f.taskList, nil
}

func (f *FakeStore) Save(taskList *task.TaskList) error {
	if taskList == nil {
		return fmt.Errorf("taskList pointer cannot be nil")
	}
	f.SaveCalls++
	f.taskList = taskList
	return nil
}

func NewFakeStore(filename string) *FakeStore {
	return &FakeStore{Filename: filename}
}

func TestSaveLoadTasks(t *testing.T) {

	t.Run("save then load preserves task data", func(t *testing.T) {
		store := NewFakeStore("dummy.json")
		list := task.NewTaskList()
		originalTask, _ := list.AddTask("buy milk")

		err := store.Save(list)

		if err != nil {
			t.Fatalf("Save failed: %v", err)
		}
		if store.SaveCalls != 1 {
			t.Error("did not make any saved calls")
		}

		taskList, err := store.Load()

		if err != nil {
			t.Fatalf("Load failed: %v", err)
		}

		if store.LoadCalls != 1 {
			t.Error("Load was not called")
		}

		// Test the actual data
		if len(taskList.Tasks) != 1 {
			t.Errorf("Number of Tasks: got %d, wanted %d", len(taskList.Tasks), 1)
		}

		if originalTask.ID != taskList.Tasks[0].ID {
			t.Errorf("ID: got %d, want %d", taskList.Tasks[0].ID, originalTask.ID)
		}

	})

	t.Run("cannot save nil value for taskList", func(t *testing.T) {
		store := NewFakeStore("test.json")

		err := store.Save(nil)

		if err == nil {
			t.Fatal("should error from nil value being saved")
		}
	})
}

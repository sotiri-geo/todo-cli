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

	t.Run("save then load task", func(t *testing.T) {
		store := NewFakeStore("dummy.json")
		list := task.NewTaskList()
		list.AddTask("buy milk")

		// persist task list
		store.Save(list)

		if store.SaveCalls != 1 {
			t.Error("did not make any saved calls")
		}

		taskList, err := store.Load()

		if store.LoadCalls != 1 {
			t.Error("Load was not called")
		}

		if err != nil {
			t.Fatal("should not error")
		}

		if len(taskList.Tasks) == 0 {
			t.Error("did not load any tasks")
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

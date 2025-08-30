package task

import (
	"slices"
	"testing"
	"time"
)

var (
	mockCreatedAt = time.Now()
)

type SpyStore struct {
	filename string
}

func (s *SpyStore) GetAll() TaskList {
	// Mocking the loading
	task1 := NewTask("add milk", true, mockCreatedAt)
	task2 := NewTask("add bread", true, mockCreatedAt)
	return TaskList{Tasks: []Task{task1, task2}}
}

func TestStore(t *testing.T) {
	t.Run("Loads correct task list from json", func(t *testing.T) {
		spyStore := SpyStore{"example.json"}
		got := LoadTasks(&spyStore)
		want := TaskList{Tasks: []Task{NewTask("add milk", true, mockCreatedAt), NewTask("add bread", true, mockCreatedAt)}}

		if !slices.Equal(got.Tasks, want.Tasks) {
			t.Errorf("got %+v, want %+v", got.Tasks, want.Tasks)
		}
	})
}

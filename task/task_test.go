package task

import (
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {

	t.Run("first task", func(t *testing.T) {
		description := "Buy milk"
		createdAt := time.Now()
		got := NewTask(description, false, createdAt)

		want := Task{ID: 0, Description: description, Completed: false, CreatedAt: createdAt}

		if got != want {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})

	// t.Run("second task task", func(t *testing.T) {
	// 	description := "Buy bread"
	// 	createdAt := time.Now()
	// 	got := NewTask(description, false, createdAt)

	// 	// ID should autoincrement
	// 	want := Task{ID: 1, Description: description, Completed: false, CreatedAt: createdAt}

	// 	if got != want {
	// 		t.Errorf("got %+v, want %+v", got, want)
	// 	}
	// })
}

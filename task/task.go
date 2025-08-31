package task

import "time"

type Task struct {
	ID          int
	Description string
	Completed   bool
	CreatedAt   time.Time
}

// Have constructors alway return pointers to the structs they create
func NewTask(description string, id int, createdAt time.Time) *Task {
	return &Task{ID: id, Description: description, Completed: false, CreatedAt: createdAt}
}

package task

import "time"

type Task struct {
	ID          int
	Description string
	Completed   bool
	CreatedAt   time.Time
}

func NewTask(description string, completed bool, createdAt time.Time) Task {
	return Task{ID: 0, Description: description, Completed: false, CreatedAt: createdAt}
}

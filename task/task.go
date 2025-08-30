package task

import "time"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

func NewTask(description string, completed bool, createdAt time.Time) Task {
	return Task{ID: 0, Description: description, Completed: false, CreatedAt: createdAt}
}

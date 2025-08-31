package task

import "time"

type Task struct {
	ID          int
	Description string
	Completed   bool
	CreatedAt   time.Time
}

type TaskList struct {
	Tasks []Task
}

func (t *Task) MarkDone() {
	t.Completed = true
}

// Have constructors alway return pointers to the structs they create
func NewTask(description string, id int, createdAt time.Time) *Task {
	return &Task{ID: id, Description: description, Completed: false, CreatedAt: createdAt}
}

func (t *TaskList) AddTask(description string) {
	task := NewTask(description, len(t.Tasks)+1, time.Now())
	t.Tasks = append(t.Tasks, *task)
}

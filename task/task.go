package task

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrEmptyTaskDescription = errors.New("description must not be empty")
)

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
func NewTask(description string, id int, createdAt time.Time) (*Task, error) {
	if len(description) == 0 {
		return &Task{}, ErrEmptyTaskDescription
	}
	return &Task{ID: id, Description: description, Completed: false, CreatedAt: createdAt}, nil
}

func (t *TaskList) AddTask(description string) (*Task, error) {
	newId := len(t.Tasks) + 1
	task, err := NewTask(description, newId, time.Now())
	if err != nil {
		return &Task{}, fmt.Errorf("Failed to add task: %w", err)
	}
	t.Tasks = append(t.Tasks, *task)
	return task, nil
}

func (t *TaskList) DeleteTask(id int) error {
	var indexToRemove int
	found := false

	for idx, task := range t.Tasks {
		if task.ID == id {
			indexToRemove = idx
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("taskId %d not found", id)
	}

	t.Tasks = append(t.Tasks[:indexToRemove], t.Tasks[indexToRemove+1:]...)

	return nil
}

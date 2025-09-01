package task

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrEmptyTaskDescription = errors.New("description must not be empty")
	ErrNotFoundTask         = errors.New("Task does not exist in task list")
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskList struct {
	Tasks  []*Task // Safer to store pointers when we need to mutate state
	nextId int
}

func (t *Task) Complete() {
	t.Completed = true
}

// Have constructors alway return pointers to the structs they create
func NewTask(description string, id int, createdAt time.Time) (*Task, error) {
	if strings.TrimSpace(description) == "" {
		return &Task{}, ErrEmptyTaskDescription
	}
	return &Task{ID: id, Description: description, Completed: false, CreatedAt: createdAt}, nil
}

func NewTaskList() *TaskList {
	return &TaskList{
		Tasks:  make([]*Task, 0),
		nextId: 0,
	}
}

func (t *TaskList) AddTask(description string) (*Task, error) {
	t.nextId++
	task, err := NewTask(description, t.nextId, time.Now())
	if err != nil {
		return &Task{}, fmt.Errorf("Failed to add task: %w", err)
	}
	t.Tasks = append(t.Tasks, task) // storing values makes a copy
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

func (t *TaskList) FindCompleted() []*Task {
	completed := make([]*Task, 0)

	for _, task := range t.Tasks {
		if task.Completed {
			completed = append(completed, task)
		}
	}

	return completed
}

func (t *TaskList) GetTask(id int) (*Task, error) {
	for _, task := range t.Tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return &Task{}, ErrNotFoundTask
}

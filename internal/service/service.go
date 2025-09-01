package service

import (
	"fmt"

	"github.com/sotiri-geo/todo-cli/internal/storage"
	"github.com/sotiri-geo/todo-cli/internal/task"
)

type TaskService struct {
	store storage.Store
}

func NewTaskService(store storage.Store) *TaskService {
	return &TaskService{store: store}
}

func (t *TaskService) AddTask(description string) (*task.Task, error) {
	// Load from storage
	taskList, err := t.store.Load()

	// Gracefully handle NotFound by creating a new taskList
	if err != nil {
		taskList = task.NewTaskList()
	}
	task, err := taskList.AddTask(description)

	t.store.Save(taskList)
	return task, err
}

func (t *TaskService) ListTasks() (*task.TaskList, error) {
	return t.store.Load()
}

func (t *TaskService) ListCompletedTasks() (*task.TaskList, error) {
	loadedTaskList, err := t.store.Load()

	if err != nil {
		return loadedTaskList, err
	}
	// filter for completed
	newTaskList := task.NewTaskList()
	newTaskList.Tasks = loadedTaskList.FindCompleted()
	return newTaskList, nil
}

func (t *TaskService) MarkCompleted(id int) (*task.Task, error) {
	loadedTaskList, errLoad := t.store.Load()

	if errLoad != nil {
		return &task.Task{}, fmt.Errorf("Failed to load task: Found %v", errLoad)
	}

	completedTask, errMark := loadedTaskList.MarkCompleted(id)

	if errMark != nil {
		return completedTask, fmt.Errorf("Failed to mark task %d id as completed: Found %v", id, errMark)
	}
	return completedTask, nil
}

func (t *TaskService) GetTask(id int) (*task.Task, error) {
	loadedTaskList, errLoad := t.store.Load()

	if errLoad != nil {
		return nil, fmt.Errorf("Failed to get task id %d: Found %w", id, errLoad)
	}

	task, errGet := loadedTaskList.GetTask(id)

	return task, errGet
}

func (t *TaskService) DeleteTask(id int) error {
	loadedTaskList, errLoad := t.store.Load()

	if errLoad != nil {
		return fmt.Errorf("Failed to delete task: %w", errLoad)
	}

	errDelete := loadedTaskList.DeleteTask(id)

	// persist back to store
	t.store.Save(loadedTaskList)
	return errDelete
}

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
	taskList := task.NewTaskList()
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

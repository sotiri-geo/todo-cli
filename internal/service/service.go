package task

import (
	"github.com/sotiri-geo/todo-cli/internal/storage"
	"github.com/sotiri-geo/todo-cli/internal/task"
)

type TaskService struct {
	store storage.Store
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

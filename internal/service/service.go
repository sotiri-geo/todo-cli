package task

import (
	"github.com/sotiri-geo/todo-cli/internal/storage"
	"github.com/sotiri-geo/todo-cli/internal/task"
)

type TaskService struct {
	store storage.Store
}

func (t *TaskService) AddTask(description string) (*task.TaskList, error) {
	taskList := task.NewTaskList()
	taskList.AddTask(description)
	t.store.Save(taskList)
	return taskList, nil
}

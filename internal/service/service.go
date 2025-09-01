package task

import (
	"fmt"

	"github.com/sotiri-geo/todo-cli/internal/storage"
	"github.com/sotiri-geo/todo-cli/internal/task"
)

type TaskService struct {
	store storage.Store
}

func (t *TaskService) AddTask(description string) (*task.TaskList, error) {
	taskList := task.NewTaskList()
	_, err := taskList.AddTask(description)
	if err != nil {
		return taskList, fmt.Errorf("Failed to add task: %v", err)
	}
	t.store.Save(taskList)
	return taskList, nil
}

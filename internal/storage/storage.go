package storage

import (
	"github.com/sotiri-geo/todo-cli/internal/task"
)

type Store interface {
	Load() (*task.TaskList, error)
	Save(*task.TaskList) error
}

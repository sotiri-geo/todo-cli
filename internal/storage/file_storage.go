package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sotiri-geo/todo-cli/internal/task"
)

type FileStore struct {
	filename string
	taskList *task.TaskList
}

func NewFileStore(filename string) *FileStore {
	return &FileStore{filename: filename}
}

func (f *FileStore) Save(taskList *task.TaskList) error {
	b, err := json.MarshalIndent(taskList, "", "  ")

	if err != nil {
		return fmt.Errorf("Failed to serialize: %v", err)
	}

	// Write to disk
	err = os.WriteFile(f.filename, b, 0644)

	if err != nil {
		return fmt.Errorf("Failed to write file: %v", err)
	}

	return nil
}

func (f *FileStore) Load() (*task.TaskList, error) {
	content, err := os.ReadFile(f.filename)

	if err != nil {
		return task.NewTaskList(), fmt.Errorf("Failed to load: %v", err)
	}

	// deserialise json
	var loadedList task.TaskList
	err = json.Unmarshal(content, &loadedList)

	if err != nil {
		return task.NewTaskList(), fmt.Errorf("Failed to deserialise: %v", err)
	}

	return &loadedList, nil
}

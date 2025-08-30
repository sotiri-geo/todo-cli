package task

type TaskStore interface {
	GetAll() TaskList
}

func LoadTasks(store TaskStore) TaskList {
	tasks := store.GetAll()
	return tasks
}

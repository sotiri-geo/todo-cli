package cmd

import (
	"fmt"

	"github.com/sotiri-geo/todo-cli/internal/task"
)

// Format outputs

func formatAdd(task task.Task) {
	output := fmt.Sprintf("✅ Task added: %q, Completed: %v", task.Description, task.Completed)
	fmt.Println(output)
}

func formatTaskRow(task task.Task) {
	var output string

	if task.Completed {
		output = fmt.Sprintf("[✓] %d. %q", task.ID, task.Description)
	} else {
		output = fmt.Sprintf("[ ] %d. %q", task.ID, task.Description)
	}
	fmt.Println(output)
}

func formatList(taskList task.TaskList) {
	fmt.Println("<<< 📝 Current tasks >>>")
	for _, task := range taskList.Tasks {
		formatTaskRow(*task)
	}
}

func formatListCompleted(taskList task.TaskList) {
	for _, task := range taskList.Tasks {
		if task.Completed {
			formatTaskRow(*task)
		}
	}
}

func formatCompleted(task task.Task) {
	output := fmt.Sprintf("👍 Task %q - Completed", task.Description)
	fmt.Println(output)
}

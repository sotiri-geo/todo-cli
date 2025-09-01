package cmd

import (
	"github.com/sotiri-geo/todo-cli/internal/service"
	"github.com/sotiri-geo/todo-cli/internal/storage"
	"github.com/spf13/cobra"
)

var completed bool
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		store := storage.NewFileStore(filename)
		svc := service.NewTaskService(store)
		if completed {
			completedList, err := svc.ListCompletedTasks()
			if err != nil {
				return err
			}
			formatList(*completedList)
		} else {
			taskList, err := svc.ListTasks()
			if err != nil {
				return err
			}
			formatList(*taskList)
		}
		return nil
	},
}

func init() {
	listCmd.Flags().BoolVarP(&completed, "completed", "c", false, "List only completed tasks.")
	rootCmd.AddCommand(listCmd)
}

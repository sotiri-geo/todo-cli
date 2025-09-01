package cmd

import (
	"fmt"

	"github.com/sotiri-geo/todo-cli/internal/service"
	"github.com/sotiri-geo/todo-cli/internal/storage"
	"github.com/spf13/cobra"
)

var completed bool
var pending bool
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		if pending && completed {
			return fmt.Errorf("flags --completed and --pending cannot be used together")
		}
		store := storage.NewFileStore(filename)
		svc := service.NewTaskService(store)

		if completed {
			completedList, err := svc.ListCompletedTasks()
			if err != nil {
				return err
			}
			fmt.Println("✅ Showing only completed tasks...")
			formatList(*completedList, formatCompleted)
		} else if pending {
			pendingList, err := svc.ListPendingTasks()
			if err != nil {
				return err
			}
			fmt.Println("📝 Showing only pending tasks...")
			formatList(*pendingList, formatTaskRow)
		} else {
			allList, err := svc.ListTasks()
			if err != nil {
				return err
			}
			fmt.Println("📋 Showing all tasks...")
			formatList(*allList, formatTaskRow)
		}
		return nil
	},
}

func init() {
	listCmd.Flags().BoolVarP(&completed, "completed", "c", false, "List only completed tasks.")
	listCmd.Flags().BoolVarP(&pending, "pending", "p", false, "List only pending tasks.")
	rootCmd.AddCommand(listCmd)
}

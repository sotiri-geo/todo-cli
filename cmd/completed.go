package cmd

import (
	"fmt"
	"strconv"

	"github.com/sotiri-geo/todo-cli/internal/service"
	"github.com/sotiri-geo/todo-cli/internal/storage"
	"github.com/spf13/cobra"
)

var completedCmd = &cobra.Command{
	Use:   "done [task id]",
	Short: "Mark a task as completed.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("Cannot parse value %q as an integer: Found %v", args[0], err)
		}
		store := storage.NewFileStore(filename)
		svc := service.NewTaskService(store)
		taskCompleted, err := svc.MarkCompleted(id)

		if err != nil {
			return err
		}
		formatCompleted(*taskCompleted)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(completedCmd)
}

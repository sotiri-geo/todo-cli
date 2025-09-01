package cmd

import (
	"fmt"
	"strconv"

	"github.com/sotiri-geo/todo-cli/internal/service"
	"github.com/sotiri-geo/todo-cli/internal/storage"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task id]",
	Short: "Delete a task from list.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		taskId, errParseInt := strconv.Atoi(args[0])
		if errParseInt != nil {
			return fmt.Errorf("Could not parse %q as a number", args[0])
		}

		store := storage.NewFileStore(filename)
		svc := service.NewTaskService(store)
		errDelete := svc.DeleteTask(taskId)

		if errDelete != nil {
			return errDelete
		}

		formatDeleted(taskId)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

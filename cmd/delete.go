package cmd

import "github.com/spf13/cobra"

var deleteCmd = &cobra.Command{
	Use:   "delete [task id]",
	Short: "Delete a task from list.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

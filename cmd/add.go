package cmd

import "github.com/spf13/cobra"

var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a new task to your todo list",
	Args:  cobra.ExactArgs(1), // Limit to 1 user input
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

// Associate it to the root (DS like a tree)
func init() {
	rootCmd.AddCommand(addCmd)
}

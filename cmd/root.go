package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const filename = "task.json"

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A simple CLI tool to manage tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Todo CLI! Use `todo --help` to see available commands.")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

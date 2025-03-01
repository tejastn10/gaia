package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tejastn10/gaia/tasks"
)

var defaultMessage = "feat: initialize project"

var rootCmd = &cobra.Command{
	Use:   "gaia [commit message]",
	Short: "Modify the first commit in a git repository",
	Long: `A CLI tool that checks if the current directory is a git repository
and updates the very first commit with a new commit message.

If no commit message is provided, a default message "feat: initialize project" will be used.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create context with cancellation
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		message := defaultMessage
		if len(args) > 0 && args[0] != "" {
			message = args[0]
		}

		err := tasks.UpdateFirstCommit(ctx, message)
		if err != nil {
			return fmt.Errorf("failed to update the first commit: %w", err)
		}

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

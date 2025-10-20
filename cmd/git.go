package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	pushFlag     bool
	distinctFlag bool
)

var gitCmd = &cobra.Command{
	Use:   "git [operation]",
	Short: "Manage Git operations within DevPilot",
	Long: `Perform various Git operations such as syncing, committing, or pushing 
with optional flags for automation.`,
	Args: cobra.ExactArgs(1), // only expect the operation name
	Run: func(cmd *cobra.Command, args []string) {
		operation := args[0]

		fmt.Printf("🚀 Running git operation: %s\n", operation)
		fmt.Printf("  --push: %v\n", pushFlag)
		fmt.Printf("  --distinct: %v\n", distinctFlag)

		// Example logic
		if operation == "sync" {
			if distinctFlag {
				fmt.Println("🔹 Committing files by folder (distinct mode)")
			}
			if pushFlag {
				fmt.Println("⬆️  Pushing changes to remote")
			}
		}
	},
}

func init() {
	gitCmd.Flags().BoolVar(&pushFlag, "push", false, "Push commits to remote after syncing")
	gitCmd.Flags().BoolVar(&distinctFlag, "distinct", false, "Group commits by folder before committing")
}

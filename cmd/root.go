// cmd/root.go
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "devpilot",
	Short: "DevPilot is your developer productivity assistant",
	Long: `DevPilot helps automate common developer workflows such as 
project scaffolding, git sync, docker management, and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to DevPilot 🚀")
		fmt.Println("Use `devpilot --help` to see available commands.")
	},
}

// Execute is the entry point for the CLI
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Register subcommands here
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(gitCmd)
}

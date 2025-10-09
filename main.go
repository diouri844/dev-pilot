package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	// Define the root command
	var rootCmd = &cobra.Command{
		Use:   "devpilot",
		Short: "DevPilot is a developer productivity CLI tool",
		Long: `DevPilot helps developers automate daily tasks like 
		project setup, git sync, docker management, and more.`,
		Run: func(cmd *cobra.Command, args []string) {
			// This runs when no subcommands are provided
			fmt.Println("Welcome to DevPilot 🚀")
			fmt.Println("Use `devpilot --help` to see available commands.")
		},
	}

	// Example of a simple subcommand
	var helloCmd = &cobra.Command{
		Use:   "hello",
		Short: "Prints a friendly greeting",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("👋 Hello, developer! DevPilot is up and running!")
		},
	}
	// Add subcommands to the root command
	rootCmd.AddCommand(helloCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

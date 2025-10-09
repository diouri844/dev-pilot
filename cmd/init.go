package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [project-path]",
	Short: "Initialize DevPilot in the specified project directory",
	Long: `Initialize DevPilot in an existing or new project directory by creating a .history file 
that tracks configuration, git, docker, and other development state.`,
	Args: cobra.ExactArgs(1), // Expect exactly one argument
	Run: func(cmd *cobra.Command, args []string) {
		projectPath := args[0]

		// Get absolute path
		absPath, err := filepath.Abs(projectPath)
		if err != nil {
			fmt.Println("❌ Invalid project path:", err)
			return
		}

		// Check if directory exists
		info, err := os.Stat(absPath)
		if os.IsNotExist(err) {
			// Directory doesn't exist → create it
			if err := os.MkdirAll(absPath, 0755); err != nil {
				fmt.Println("❌ Failed to create project folder:", err)
				return
			}
			fmt.Printf("📁 Created new project folder: %s\n", absPath)
		} else if err != nil {
			fmt.Println("❌ Error accessing path:", err)
			return
		} else if !info.IsDir() {
			fmt.Printf("❌ The path '%s' exists but is not a directory.\n", absPath)
			return
		}

		// Path to .history file
		historyPath := filepath.Join(absPath, ".history")

		// Check if already initialized
		if initialized, err := AlreadyInitialized(historyPath); initialized {
			fmt.Println("❌", err)
			return
		} else if err != nil {
			fmt.Println("⚠️", err)
			return
		}

		// Prepare history data
		projectName := filepath.Base(absPath)
		history := ProjectHistory{
			ProjectName: projectName,
			CreatedAt:   time.Now(),
			LastCommand: "init",
		}

		// Save .history file
		if err := SaveHistory(historyPath, history); err != nil {
			fmt.Println("❌ Failed to write .history file:", err)
			return
		}

		fmt.Printf("✅ DevPilot initialized for project '%s'\n", projectName)
		fmt.Printf("📝 .history file created at '%s'\n", historyPath)
	},
}

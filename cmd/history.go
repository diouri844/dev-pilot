// cmd/history.go
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// ProjectHistory represents the stored state of a DevPilot project
type ProjectHistory struct {
	ProjectName string    `json:"project_name"`
	CreatedAt   time.Time `json:"created_at"`
	Git         struct {
		RemoteURL     string `json:"remote_url,omitempty"`
		CurrentBranch string `json:"current_branch,omitempty"`
	} `json:"git,omitempty"`
	Docker struct {
		Containers []string `json:"containers,omitempty"`
		Images     []string `json:"images,omitempty"`
	} `json:"docker,omitempty"`
	LastCommand string `json:"last_command"`
}

// SaveHistory writes the ProjectHistory to a .history file in the given path
func SaveHistory(
	path string,
	history ProjectHistory) error {
	data, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to serialize history: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write .history file: %w", err)
	}

	return nil
}

// LoadHistory reads an existing .history file and returns its data
func LoadHistory(path string) (*ProjectHistory, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read .history file: %w", err)
	}

	var history ProjectHistory
	if err := json.Unmarshal(data, &history); err != nil {
		return nil, fmt.Errorf("failed to parse .history file: %w", err)
	}
	return &history, nil
}

// add new function taht check if the history already exist (project already inisialized )
func AlreadyInitialized(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		// File exists
		return true, fmt.Errorf("project already initialized 🙂")
	}
	if os.IsNotExist(err) {
		// File does not exist → not initialized yet
		return false, nil
	}
	// Some other error (e.g., permission issue)
	return false, fmt.Errorf("failed to check initialization: %w", err)
}

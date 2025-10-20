// cmd/config.go
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	ProjectPath string `json:"projectPath"`
	RepoURL     string `json:"repoUrl,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
}

var GlobalConfig *Config

// LoadConfig loads the .history file for the given path
func LoadConfig(projectPath string) error {
	historyPath := filepath.Join(projectPath, ".history")
	data, err := os.ReadFile(historyPath)
	if err != nil {
		return fmt.Errorf("failed to read .history: %v", err)
	}

	cfg := &Config{}
	if err := json.Unmarshal(data, cfg); err != nil {
		return fmt.Errorf("failed to parse .history: %v", err)
	}

	cfg.ProjectPath = projectPath
	GlobalConfig = cfg
	return nil
}

// MustGetConfig returns the current config or exits if missing
func MustGetConfig() *Config {
	if GlobalConfig == nil {
		fmt.Println("❌ No project initialized. Run `devpilot init <path>` first.")
		os.Exit(1)
	}
	return GlobalConfig
}

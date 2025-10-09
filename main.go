// main.go file
package main

import (
	"fmt"
	"os"

	"devpilot/cmd"
)

func main() {
	// Execute the root command
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

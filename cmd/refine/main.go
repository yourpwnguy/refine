package main

import (
	"github.com/iaakanshff/refine/pkg/runner"
)

func main() {
	// Handle version and help arguments
	runner.HandleVersionAndHelp()

	// Check if input is from stdin
	stdin := runner.IsInputFromStdin()

	// If data is from standard input
	if stdin {
		runner.HandleStdinInput()

	// If data is directly from a file
	} else {
		runner.HandleFileInput()
	}
}

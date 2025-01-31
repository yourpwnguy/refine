package main

import (
	"github.com/yourpwnguy/refine/pkg/middleflow"
	"github.com/yourpwnguy/refine/pkg/utils"
)

func main() {

	// Handle version and help arguments
	utils.HandleVersionAndHelp()

	// Check if input is from stdin
	stdin := utils.IsInputFromStdin()

	// If data is from standard input
	if stdin {
		middleflow.HandleStdinInput()

		// If data is directly from a file
	} else {
		middleflow.HandleFileInput()
	}
}

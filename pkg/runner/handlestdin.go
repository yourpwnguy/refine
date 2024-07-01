package runner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Handling input from the standard input
func HandleStdin(fflag bool) (*map[string]struct{}, int, int) {

	// Read the filename from stdin
	temp := make(map[string]struct{})
	scanner := bufio.NewScanner(os.Stdin)

	// Counter
	var totalLinesCount int = 0

	// Storing all the lines into the Map from the file
	for scanner.Scan() {
		if line := strings.TrimSpace(scanner.Text()); line != "" {
			temp[line] = struct{}{}
			totalLinesCount++
		}
	}

	// Print the unique lines to stdout
	if !fflag {
		for _, sorted := range *SortMap(&temp) {
			fmt.Println(sorted)
		}
	}
	return &temp, len(temp), totalLinesCount
}

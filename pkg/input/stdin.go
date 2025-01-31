package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/yourpwnguy/refine/pkg/utils"
)

// Handling input from the standard input
func HandleStdin(fflag bool) (map[string]struct{}, int, int, time.Time) {

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

	// Start the timer after processing stdin to exclude delays caused by external applications providing data.
	startTime := time.Now()

	// Print the unique lines to stdout
	if !fflag {
		for _, sorted := range utils.SortMap(temp) {
			fmt.Println(sorted)
		}
	}
	return temp, len(temp), totalLinesCount, startTime
}

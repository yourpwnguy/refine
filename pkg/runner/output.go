package runner

import (
	"fmt"
	"os"
)

// ProcessOutput truncates the file and writes new data from temp map
func processOutput(fn string, temp *map[string]struct{}) error {

	// Open the file with write-only and truncate flags to clear its contents
	f, err := os.OpenFile(fn, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf(Errfix+" Error opening file for writing: %v", err)
	}
	defer f.Close()

	// Move the file cursor to the beginning of file for writing
	_, err = f.Seek(0, 0)
	if err != nil {
		return fmt.Errorf(Errfix+" Error seeking file:", err)
	}

	// Call SortMap once to get the sorted slice
	sortedLines := SortMap(temp)

	// Iterate over the unique lines and write them to the file
	for _, line := range sortedLines {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf(Errfix+" Error writing to file:", err)
		}
	}
	return nil
}

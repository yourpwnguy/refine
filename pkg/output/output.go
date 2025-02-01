package output

import (
	"fmt"
	"os"
	"strings"

	"github.com/yourpwnguy/refine/pkg/common"
)

// ProcessOutput writes new data from file map
func ProcessOutput(f *os.File, temp map[string]struct{}) error {

	// Move the file cursor to the beginning of file for writing
	_, err := f.Seek(0, 0)
	if err != nil {
		return fmt.Errorf(common.Errfix+" Error seeking file:", err)
	}

	// Call SortMap once to get the sorted slice
	sortedLines := common.SortMap(temp)

	// Join all lines with newline and write them in one shot
	_, err = f.WriteString(strings.Join(sortedLines, "\n") + "\n")
	if err != nil {
		return fmt.Errorf(common.Errfix+" Error writing to file:", err)
	}

	return nil
}

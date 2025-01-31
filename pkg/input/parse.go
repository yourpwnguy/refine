package input

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/yourpwnguy/refine/pkg/common"
)

// parseInput splits a comma-separated string of file names into a slice of strings.
func ParseInput(files string) []string {
	return strings.Split(files, ",")
}

// readFiles reads and returns the list of files in the specified directory.
func ReadFiles(dir string) []fs.DirEntry {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, common.Errfix, "Error fetching all the files")
		os.Exit(0)
	}
	return files
}

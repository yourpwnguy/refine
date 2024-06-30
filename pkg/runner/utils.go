package runner

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

// isDir checks if the given path is a directory.
func isDir(dir string) bool {
	info, _ := os.Stat(dir)
	return info.IsDir()
}

// readFiles reads and returns the list of files in the specified directory.
func readFiles(dir string) []fs.DirEntry {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, Errfix, "Error fetching all the files")
		os.Exit(0)
	}
	return files
}

// parseInput splits a comma-separated string of file names into a slice of strings.
func parseInput(files string) []string {
	return strings.Split(files, ",")
}

// HandleVersionAndHelp processes the command-line arguments to check for version and help flags.
func HandleVersionAndHelp() {
	for _, arg := range os.Args[1:] {
		if arg == "-v" || arg == "--version" {
			CheckVersion()
			os.Exit(0)
		} else if arg == "-h" || arg == "--help" {
			PrintUsage()
			os.Exit(0)
		}
	}
}

// IsInputFromStdin checks if the input is being provided from the standard input (stdin).
func IsInputFromStdin() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

// shouldSkip checks if a file should be skipped based on the exceptions list.
func shouldSkip(fileName string, exceptions []string) bool {
	for _, exception := range exceptions {
		if fileName == exception {
			return true
		}
	}
	return false
}

package runner

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// HandleStdinInput processes input coming from stdin.
func HandleStdinInput() {

	// Getting the start time
	currTime := time.Now()

	var (
		uniqueLinesCount     int = 0
		totalLinesCount      int = 0
		totalStdinLinesCount int = 0
		filename             string
		err                  error
	)

	// If there are more than 2 arguments in stdin mode, then just exit
	if len(os.Args) > 2 {
		fmt.Fprintln(os.Stderr, Errfix, "Invalid arguments or usage")
		os.Exit(0)

		// Example: refine file.txt
	} else if len(os.Args) == 2 {
		var stdinMap *map[string]struct{}

		// Getting all the lines from stdin
		stdinMap, _, totalStdinLinesCount = HandleStdin(true)

		// Getting all the lines from the given file
		uniqueLinesCount, totalLinesCount, filename, err = HandleFile(os.Args[1], stdinMap, true)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, "%s File: %s, Original: %d lines, Unique: %d lines, Processed in %v\n",
			Succfix, filepath.Base(filename), totalLinesCount+totalStdinLinesCount, uniqueLinesCount, time.Since(currTime))

		// Example: cat file.txt | refine
	} else {
		_, uniqueLinesCount, totalStdinLinesCount = HandleStdin(false)
		fmt.Fprintf(os.Stderr, "%s Original: %d lines, Unique: %d lines, Processed in %v\n",
			Succfix, totalLinesCount+totalStdinLinesCount, uniqueLinesCount, time.Since(currTime))
	}

}

// HandleFileInput processes file inputs based on command-line arguments.
func HandleFileInput() {

	// Getting the start time
	currTime := time.Now()

	var (
		uniqueLinesCount int = 0
		totalLinesCount  int = 0
		filename         string
		err              error
	)

	switch len(os.Args) {

	case 1:
		PrintUsage()
		os.Exit(0)
	// Example: refine file.txt
	case 2:
		uniqueLinesCount, totalLinesCount, filename, err = HandleFile(os.Args[1], nil, false)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, "%s File: %s, Original: %d lines, Unique: %d lines, Processed in %v\n",
			Succfix, filepath.Base(filename), totalLinesCount, uniqueLinesCount, time.Since(currTime))

	case 3:

		// Checking if the wildcard is given
		if os.Args[1] == "-w" || os.Args[1] == "--wildcard" {
			HandleWildcard(os.Args[2], nil)

			// Checking if multiple files are given max: 2
		} else {
			uniqueLinesCount, totalLinesCount, filename, err = HandleMultipleFiles(os.Args[1], os.Args[2])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(0)
			}
			fmt.Fprintf(os.Stderr, "%s File: %s, Original: %d lines, Unique: %d lines, Processed in %v\n",
				Succfix, filepath.Base(filename), totalLinesCount, uniqueLinesCount, time.Since(currTime))
		}

	case 5:

		// Checking if wildcard is given with --we ( wildcard exception )
		if os.Args[1] == "-w" || os.Args[1] == "--wildcard" && os.Args[3] == "-we" || os.Args[3] == "--wildcard-exception" {
			exceptions := parseInput(os.Args[4])
			HandleWildcard(os.Args[2], exceptions)
		} else {
			fmt.Fprintln(os.Stderr, Errfix, "Invalid arguments or usage")
			os.Exit(0)
		}

	// Otherwise incorrect usage
	default:
		fmt.Fprintln(os.Stderr, Errfix, "Invalid arguments or usage")
		os.Exit(0)
	}
}

// HandleWildcard processes files in a directory, optionally excluding specified exceptions.
func HandleWildcard(dir string, exceptions []string) {

	// Getting the absolute directory path
	absDirname, err := filepath.Abs(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting absolute path:", err)
		os.Exit(0)
	}

	// Checking if it's really a directory
	if isDir(absDirname) {

		// readFiles return the files in the directory
		for _, file := range readFiles(absDirname) {

			// Processing time for each iteration
			currTime := time.Now()

			// Checking if it's really a text file
			if file.Type().IsRegular() {

				// Comparing it against the list of exceptions files so that if match we can skip the file
				if shouldSkip(file.Name(), exceptions) {
					continue
				}

				// Joining the dir and file to create an absolute path to the file so that file with the same name in the current directory doesn't get processed
				absFilePath := filepath.Join(absDirname, file.Name())
				uniqueLinesCount, totalLinesCount, filename, err := HandleFile(absFilePath, nil, false)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}
				fmt.Fprintf(os.Stderr, "%s File: %s, Original: %d lines, Unique: %d lines, Processed in %v\n",
					Succfix, filepath.Base(filename), totalLinesCount, uniqueLinesCount, time.Since(currTime))
			}
		}
	} else {
		fmt.Fprintln(os.Stderr, Errfix, dir, "is not a directory.")
		os.Exit(0)
	}
}

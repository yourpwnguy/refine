package middleflow

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/yourpwnguy/refine/pkg/common"
	"github.com/yourpwnguy/refine/pkg/input"
	"github.com/yourpwnguy/refine/pkg/output"
	"github.com/yourpwnguy/refine/pkg/types"
	"github.com/yourpwnguy/refine/pkg/utils"
)

// HandleStdinInput processes input coming from stdin.
func HandleStdinInput() {

	var (
		uniqueLinesCount     int = 0
		totalLinesCount      int = 0
		totalStdinLinesCount int = 0
		err                  error
		stdinMap             map[string]struct{}
		startTime            time.Time
	)

	// If there are more than 2 arguments in stdin mode, then just exit
	if len(os.Args) > 2 {
		fmt.Fprintln(os.Stderr, common.Errfix, "Invalid arguments or usage")
		os.Exit(1)

		// Example: cat x.txt | refine file.txt
	} else if len(os.Args) == 2 {

		// Getting all the lines from stdin
		stdinMap, _, totalStdinLinesCount, startTime = input.HandleStdin(true)

		// Getting all the lines from the given file
		uniqueLinesCount, totalLinesCount, err = input.HandleFile(os.Args[1], stdinMap, true)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		output.BeautifyPrint(types.Params{
			Filename:         "",
			OutputFile:       os.Args[1],
			TotalLinesCount:  totalLinesCount + totalStdinLinesCount,
			UniqueLinesCount: uniqueLinesCount,
			TimeTaken:        time.Since(startTime),
		})

		// Example: cat file.txt | refine
	} else {
		_, uniqueLinesCount, totalStdinLinesCount, startTime = input.HandleStdin(false)
		output.BeautifyPrint(types.Params{
			Filename:         "",
			OutputFile:       "",
			TotalLinesCount:  totalLinesCount + totalStdinLinesCount,
			UniqueLinesCount: uniqueLinesCount,
			TimeTaken:        time.Since(startTime),
		})
	}

}

// HandleFileInput processes file inputs based on command-line arguments.
func HandleFileInput() {

	// Getting the start time
	currTime := time.Now()

	var (
		uniqueLinesCount int = 0
		totalLinesCount  int = 0
		err              error
	)

	switch len(os.Args) {

	case 1:
		common.PrintUsage()
		os.Exit(1)

	// Example: refine file.txt
	case 2:
		uniqueLinesCount, totalLinesCount, err = input.HandleFile(os.Args[1], nil, false)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		output.BeautifyPrint(types.Params{
			Filename:         os.Args[1],
			OutputFile:       os.Args[1],
			TotalLinesCount:  totalLinesCount,
			UniqueLinesCount: uniqueLinesCount,
			TimeTaken:        time.Since(currTime),
		})

	case 3:

		// Checking if the wildcard is given
		if os.Args[1] == "-w" || os.Args[1] == "--wildcard" {
			HandleWildcard(os.Args[2], nil)

			// Checking if multiple files are given max: 2
		} else {
			uniqueLinesCount, totalLinesCount, err = input.HandleMultipleFiles(os.Args[1], os.Args[2])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			output.BeautifyPrint(types.Params{
				Filename:         os.Args[1],
				OutputFile:       os.Args[2],
				TotalLinesCount:  totalLinesCount,
				UniqueLinesCount: uniqueLinesCount,
				TimeTaken:        time.Since(currTime),
			})
		}

	case 5:

		// Checking if wildcard is given with --we ( wildcard exception )
		if os.Args[1] == "-w" || os.Args[1] == "--wildcard" && os.Args[3] == "-we" || os.Args[3] == "--wildcard-exception" {
			exceptions := input.ParseInput(os.Args[4])
			HandleWildcard(os.Args[2], exceptions)
		} else {
			fmt.Fprintln(os.Stderr, common.Errfix, "Invalid arguments or usage")
			os.Exit(1)
		}

	// Otherwise incorrect usage
	default:
		fmt.Fprintln(os.Stderr, common.Errfix, "Invalid arguments or usage")
		os.Exit(1)
	}
}

// HandleWildcard processes files in a directory, optionally excluding specified exceptions.
func HandleWildcard(dir string, exceptions []string) {

	// Getting the absolute directory path
	absDirname, err := filepath.Abs(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting absolute path:", err)
		os.Exit(1)
	}

	// Checking if it's really a directory
	if utils.IsDir(absDirname) {

		// readFiles return the files in the directory
		for _, file := range input.ReadFiles(absDirname) {

			// Processing time for each iteration
			currTime := time.Now()

			// Checking if it's really a text file
			if file.Type().IsRegular() {

				// Comparing it against the list of exceptions files so that if match we can skip the file
				if utils.ShouldSkip(file.Name(), exceptions) {
					continue
				}

				// Joining the dir and file to create an absolute path to the file so that file with the same name in the current directory doesn't get processed
				absFilePath := filepath.Join(absDirname, file.Name())
				uniqueLinesCount, totalLinesCount, err := input.HandleFile(absFilePath, nil, false)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}
				output.BeautifyPrint(types.Params{
					Filename:         file.Name(),
					OutputFile:       file.Name(),
					TotalLinesCount:  totalLinesCount,
					UniqueLinesCount: uniqueLinesCount,
					TimeTaken:        time.Since(currTime),
				})
			}
		}
	} else {
		fmt.Fprintln(os.Stderr, common.Errfix, dir, "is not a directory.")
		os.Exit(1)
	}
}

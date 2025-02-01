package middleflow

import (
	"fmt"
	"os"
	"time"

	"github.com/yourpwnguy/refine/pkg/common"
	"github.com/yourpwnguy/refine/pkg/input"
	"github.com/yourpwnguy/refine/pkg/output"
	"github.com/yourpwnguy/refine/pkg/types"
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
			OutputFile:       os.Args[1],
			TotalLinesCount:  totalLinesCount + totalStdinLinesCount,
			UniqueLinesCount: uniqueLinesCount,
			TimeTaken:        time.Since(startTime),
		})

		// Example: cat file.txt | refine
	} else {
		_, uniqueLinesCount, totalStdinLinesCount, startTime = input.HandleStdin(false)
		output.BeautifyPrint(types.Params{
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
			OutputFile:       os.Args[1],
			TotalLinesCount:  totalLinesCount,
			UniqueLinesCount: uniqueLinesCount,
			TimeTaken:        time.Since(currTime),
		})

	case 3:

		// Checking if the wildcard is given
		if os.Args[1] == "-w" || os.Args[1] == "--wildcard" {
			input.HandleWildcard(os.Args[2], nil)

			// Checking if multiple files are given max: 2
		} else {
			uniqueLinesCount, totalLinesCount, err = input.HandleMultipleFiles(os.Args[1], os.Args[2])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			output.BeautifyPrint(types.Params{
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
			input.HandleWildcard(os.Args[2], exceptions)
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

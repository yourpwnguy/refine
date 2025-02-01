package input

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/yourpwnguy/refine/pkg/common"
	"github.com/yourpwnguy/refine/pkg/output"
	"github.com/yourpwnguy/refine/pkg/types"
	"github.com/yourpwnguy/refine/pkg/utils"
)

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
		for _, file := range ReadFiles(absDirname) {

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
				uniqueLinesCount, totalLinesCount, err := HandleFile(absFilePath, nil, false)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}
				output.BeautifyPrint(types.Params{
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

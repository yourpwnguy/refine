package input

import (
	"fmt"
	"os"

	"github.com/yourpwnguy/refine/pkg/common"
	"github.com/yourpwnguy/refine/pkg/output"
	"github.com/yourpwnguy/refine/pkg/utils"
)

// Handling the given file provided with the associated map
func HandleFile(fn string, stdinMap map[string]struct{}, toBeCreated bool) (int, int, error) {

	// Check if file exists
	_, err := os.Stat(fn)
	var f *os.File

	// Checking if the file need to be created
	// Ex: cat file.txt | refine file.txt
	if toBeCreated {

		// File doesn't exist, create it
		// File exists, open with read and write permissions
		f, err = utils.OpenOrCreateFile(fn, true)
		if err != nil {
			return 0, 0, fmt.Errorf("%s Error opening file: %w", common.Errfix, err)
		}
		defer f.Close()

		// Otherwise check if it exists or not
		// Ex: refine file.txt
	} else {
		if os.IsNotExist(err) {
			return 0, 0, fmt.Errorf("%s Provided file '%s' doesn't exist: %w", common.Errfix, fn, err)
		} else {
			// File exists, open with read and write permissions
			f, err = utils.OpenOrCreateFile(fn, true)
			if err != nil {
				return 0, 0, fmt.Errorf("%s Error opening file: %w", common.Errfix, err)
			}
			defer f.Close()
		}
	}

	fileMap := make(map[string]struct{})

	totalLinesCount, err := utils.ReadLinesFromFileAndStdin(f, fileMap, stdinMap)
	if err != nil {
		return 0, 0, err
	}

	f, err = utils.OpenAndTruncate(f)
	if err != nil {
		return 0, 0, fmt.Errorf("%s Error truncating file: %w", common.Errfix, err)
	}
	defer f.Close()

	// Storing the output in the file
	if err := output.ProcessOutput(f, fileMap); err != nil {
		return 0, 0, err
	}

	return len(fileMap), totalLinesCount, nil
}

// Handling multiple files
func HandleMultipleFiles(fn1, fn2 string) (int, int, error) {

	// Creating/Opening the first file
	f1, err := utils.OpenOrCreateFile(fn1, false)
	if err != nil {
		return 0, 0, fmt.Errorf("%s Error opening file: %w", common.Errfix, err)
	}
	defer f1.Close()

	fileMap := make(map[string]struct{})
	totalLinesCount, err := utils.ReadLinesFromFileAndStdin(f1, fileMap, nil)
	if err != nil {
		return 0, 0, err
	}

	// Creating/Opening the second file for storing the output
	f2, err := utils.OpenOrCreateFile(fn2, true)
	if err != nil {
		return 0, 0, fmt.Errorf("%s Error opening file: %w", common.Errfix, err)
	}
	defer f2.Close()

	tempLinesCount, err := utils.ReadLinesFromFileAndStdin(f2, fileMap, nil)
	if err != nil {
		return 0, 0, err
	}

	f2, err = utils.OpenAndTruncate(f2)
	if err != nil {
		return 0, 0, fmt.Errorf("%s Error truncating file: %w", common.Errfix, err)
	}
	defer f2.Close()

	// Storing the output in the file 2
	if err := output.ProcessOutput(f2, fileMap); err != nil {
		return 0, 0, err
	}
	return len(fileMap), totalLinesCount + tempLinesCount, nil
}

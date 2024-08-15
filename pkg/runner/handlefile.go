package runner

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yourpwnguy/gostyle"
)

var (
	// Some colored messages ( mainly prefix )
	g       = gostyle.New()
	Succfix = "[" + g.Blue("INF") + "]"
	Errfix  = "[" + g.Red("ERR") + "]"
)

// Handling the given file provided with the associated map
func HandleFile(fn string, stdinMap *map[string]struct{}, toBeCreated bool) (int, int, string, error) {
	// Check if file exists
	_, err := os.Stat(fn)
	var f *os.File

	// Checking if the file need to be created
	// Ex: cat file.txt | refine file.txt
	if toBeCreated {
		if os.IsNotExist(err) {

			// File doesn't exist, create it
			f, err = os.Create(fn)
			if err != nil {
				return 0, 0, fn, fmt.Errorf(Errfix+" Error creating file:", err)
			}
			defer f.Close()

		} else if err != nil {

			// Other error occurred, handle accordingly
			return 0, 0, fn, fmt.Errorf(Errfix+" Error checking file status:", err)

		} else {
			// File exists, open with read and write permissions
			f, err = os.OpenFile(fn, os.O_RDWR|os.O_APPEND, 0644)
			if err != nil {
				return 0, 0, fn, fmt.Errorf(Errfix+" Error opening file:", err)
			}
			defer f.Close()
		}

		// Otherwise check if it exists or not
		// Ex: refine file.txt
	} else {
		if os.IsNotExist(err) {
			return 0, 0, fn, fmt.Errorf(Errfix + " Provided file doesn't exist: " + filepath.Base(fn))
		} else {
			// File exists, open with read and write permissions
			f, err = os.OpenFile(fn, os.O_RDWR|os.O_APPEND, 0644)
			if err != nil {
				return 0, 0, fn, fmt.Errorf(Errfix+" Error opening file:", err)
			}
			defer f.Close()
		}
	}

	// Map for storing all the lines from file
	temp := make(map[string]struct{})

	// Counter
	var totalLinesCount int = 0
	// For reading the existing lines from the file
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if line := strings.TrimSpace(scanner.Text()); line != "" {
			temp[line] = struct{}{}
			totalLinesCount++
		}
	}

	// Putting the lines from stdinMap to temp Map
	if stdinMap != nil {
		for line := range *stdinMap {
			temp[line] = struct{}{}
		}
	}

	// Storing the output in the file
	if err := processOutput(fn, &temp); err != nil {
		return 0, 0, fn, err
	}

	return len(temp), totalLinesCount, fn, nil
}

// Handling multiple files
func HandleMultipleFiles(fn1, fn2 string) (int, int, string, error) {

	// Open the file with read and write permissions
	f1, err := os.OpenFile(fn1, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return 0, 0, fn2, fmt.Errorf(Errfix+" Error opening file:", err)
	}
	defer f1.Close()

	// Map for storing all the lines from file
	temp := make(map[string]struct{})

	// Counter
	var totalLinesCount int = 0

	// For reading the existing lines from the file 1
	scanner := bufio.NewScanner(f1)
	for scanner.Scan() {
		if line := strings.TrimSpace(scanner.Text()); line != "" {
			temp[line] = struct{}{}
			totalLinesCount++
		}
	}

	// Creating the second file for storing the output
	f2, err := os.Create(fn2)
	if err != nil {
		return 0, 0, fn2, fmt.Errorf(Errfix+" Error creating file:", err)
	}
	defer f2.Close()

	// Storing the output in the file 2
	if err := processOutput(fn2, &temp); err != nil {
		return 0, 0, fn2, err
	}
	return len(temp), totalLinesCount, fn2, nil
}

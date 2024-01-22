package utilities

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

// Handling the given file provided with the associated map
func HandleFile(fileName string, stdinMap map[string]struct{}, regExp string) {

	// Open the file with read and write permissions
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Make sure to close the file when you're done with it

	temp := make(map[string]struct{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		temp[line] = struct{}{}
	}

	for key := range stdinMap {
		temp[key] = struct{}{}
	}

	// Apply regular expression filter
	if regExp != "" {
		temp = ExtractUsingRegex(temp, regExp)
		}

	// Move the file cursor to the beginning for writing
	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return
	}

	// Truncate, i.e, clear the content before writing new one
	err = file.Truncate(0)
	if err != nil {
		fmt.Print("Error truncating file...")
	}

	// Iterate over the unique lines and write them to the file
	for key := range temp {
		_, err := file.WriteString(key + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}
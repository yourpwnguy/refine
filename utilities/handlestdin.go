package utilities

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Handling input from the standard input
func HandleStdin(fileflag bool, regExp string) map[string]struct{} {

	// Read the filename from stdin
	temp := make(map[string]struct{})
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			words := strings.Fields(line)
			for _, word := range words {
				temp[word] = struct{}{}
			}
		}
	}

	// Print the unique lines to stdout
	if !fileflag {
		if regExp != "" {
			temp = ExtractUsingRegex(temp, regExp)
			for extractedPart := range temp {
				fmt.Println(extractedPart)
			}
		} else {
			for key := range temp {
				fmt.Println(key)
			}
		}
		return nil
	}
	return temp
}

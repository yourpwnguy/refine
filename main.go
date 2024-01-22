package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var (
		regExp   = ""
		fileFlag bool
	)

	stat, _ := os.Stdin.Stat()

	// Third option ( Reading from stdout and writing to new File )
	if (stat.Mode()&os.ModeCharDevice) == 0 && len(os.Args) > 1 {
		if len(os.Args) == 4 {
			if os.Args[1] == "-e" {
				regExp = os.Args[2]
				fileFlag = true
				stdinMap := handleStdin(fileFlag, regExp)
				handleFile(os.Args[3], stdinMap, regExp)
			} else {
				strings.Contains(os.Args[1], "-")
				fmt.Println("Unknown flag provided...")
				printUsage()
				}

		} else if len(os.Args) == 3 {
			if os.Args[1] == "-e" {
				regExp = os.Args[2]
				fileFlag = false
				handleStdin(fileFlag, regExp)
			} else {
				strings.Contains(os.Args[1], "-") 
				fmt.Println("Unknown flag provided...")
				printUsage()
			}

		} else if len(os.Args) == 2 { 
			if os.Args[1] == "-e" {
				fmt.Println("No expression is provided...")
				printUsage()
			} else if os.Args[1] == "-h" || os.Args[1] == "help" {
				printUsage()
			} else {
				fileFlag = true
				stdinMap := handleStdin(fileFlag, regExp)
				handleFile(os.Args[1], stdinMap, regExp)
			}
		} else {
			printUsage()
		}

	} else if (stat.Mode() & os.ModeCharDevice) == 0 {
		fileFlag = false
		handleStdin(fileFlag, regExp)

	} else if len(os.Args) == 4 {
		if os.Args[1] == "-e" {
			regExp = os.Args[2]
			fileFlag = true
			handleFile(os.Args[3], nil, regExp)
		} else if strings.Contains(os.Args[1], "-") {
			fmt.Println("Unknown flag provided...")
			printUsage()
		} else if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
			fmt.Println("Error: File does not exist")
			return
		} else {
			handleFile(os.Args[1], nil, regExp)
		}
	
	} else if len(os.Args) == 3 {
		if os.Args[1] == "-e" {
			fmt.Println("Expression and file is not provided...")
			printUsage()
		}
	} else if len(os.Args) == 2 {
		if os.Args[1] == "-h" || os.Args[1] == "help" {
			printUsage()
		} else if os.Args[1] == "-e" {
			fmt.Println("No expression is provided...")
			printUsage()
		} else if strings.Contains(os.Args[1], "-") {
			fmt.Println("Unknown flag provided...")
			printUsage()
		} else if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
			fmt.Println("Error: File does not exist")
			return
		} else {
			handleFile(os.Args[1], nil, regExp)
		}
	} else {
		printUsage()
	}
}

func printUsage() {
	fmt.Print("A tool to process and write unique lines written by Aakansh\n\n")
	fmt.Print("Usage: seek <filename>\n\n")
	fmt.Print("Basic Functioning: \n\n")
	fmt.Println("  seek file.txt						(Read and write to the same file)")
	fmt.Println("  cat file.txt | seek					(Read from stdin and display to the stdout)")
	fmt.Println("  cat file.txt | seek newfile.txt			(Read from stdin and write to a specific file)")
	fmt.Print("\nAdvanced Functioning: \n\n")
	fmt.Println("  seek -e 'expression'					(Regular expressions for filtering)")
}

func handleFile(fileName string, stdinMap map[string]struct{}, regExp string) {

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
		temp = extractUsingRegex(temp, regExp)
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

func handleStdin(fileflag bool, regExp string) map[string]struct{} {

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
			temp = extractUsingRegex(temp, regExp)
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

func extractUsingRegex(temp map[string]struct{}, regExp string) map[string]struct{} {
	filtered := make(map[string]struct{})
	re := regexp.MustCompile(regExp)

	for key := range temp {
		if strings.TrimSpace(key) != "" {
			matches := re.FindAllStringSubmatch(key, -1)
			for _, match := range matches {
				if len(match) > 1 {
					// Extract the full URL including the protocol
					extractedPart := match[0]
					filtered[extractedPart] = struct{}{}
				}
			}
		}
	}
	return filtered
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{
		Use:   "seek",
		Short: "A tool to process and write unique lines written by Aakansh",
		Run: func(cmd *cobra.Command, args []string) {

			stat, _ := os.Stdin.Stat()
			var fileFlag bool

			if (stat.Mode()&os.ModeCharDevice) == 0 && len(os.Args) == 2 {
				fileFlag = true
				stdinMap := handleStdin(fileFlag)
				handleFile(os.Args[1], stdinMap)
			} else if (stat.Mode() & os.ModeCharDevice) == 0 {
				fileFlag = false
				handleStdin(fileFlag)
			} else if len(os.Args) == 2 {
				handleFile(os.Args[1])
			} else {
				fmt.Print("A tool to process and write unique lines written by Aakansh\n\n")
				fmt.Print("Usage: seek <filename>\n\n")
				fmt.Println("  seek file.txt						(Read from and write to the same file)")
				fmt.Println("  cat file.txt | seek					(Read from stdin and display to the stdout)")
				fmt.Println("  cat file.txt | seek newfile.txt			(Read from stdin and write to a specific file)")
			}
		},
	}

	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Print("A tool to process and write unique lines written by Aakansh\n\n")
		fmt.Print("Usage:", "\n\n")
		fmt.Println("  seek file.txt						(Read from and write to the same file)")
		fmt.Println("  cat file.txt | seek					(Read from stdin and display to the stdout)")
		fmt.Println("  cat file.txt | seek newfile.txt			(Read from stdin and write to a specific file)")
	})
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing command...")
	}
}

func handleFile(fileName string, stdinMap ...map[string]struct{}) {
	// Open the file with read and write permissions

	temp := make(map[string]struct{})
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Make sure to close the file when you're done with it

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		temp[line] = struct{}{}
	}

	for _, tempMap := range stdinMap {
		for key := range tempMap {
			temp[key] = struct{}{}
		}
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

func handleStdin(fileflag bool) map[string]struct{} {
	// Read the filename from stdin
	temp := make(map[string]struct{})
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		temp[line] = struct{}{}
	}

	// Print the unique lines to stdout
	if !fileflag {
		for key := range temp {
			fmt.Println(key)
		}
	}
	return temp
}

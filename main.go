package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/iaakanshff/seek/utilities"
	"github.com/iaakanshff/seek/usage"
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
				stdinMap := utilities.HandleStdin(fileFlag, regExp)
				utilities.HandleFile(os.Args[3], stdinMap, regExp)
			} else {
				strings.Contains(os.Args[1], "-")
				fmt.Println("Unknown flag provided...")
				usage.PrintUsage()
			}

		} else if len(os.Args) == 3 {
			if os.Args[1] == "-e" {
				regExp = os.Args[2]
				fileFlag = false
				utilities.HandleStdin(fileFlag, regExp)
			} else {
				strings.Contains(os.Args[1], "-")
				fmt.Println("Unknown flag provided...")
				usage.PrintUsage()
			}

		} else if len(os.Args) == 2 {
			if os.Args[1] == "-e" {
				fmt.Println("No expression is provided...")
				usage.PrintUsage()
			} else if os.Args[1] == "-h" || os.Args[1] == "help" {
				usage.PrintUsage()
			} else {
				fileFlag = true
				stdinMap := utilities.HandleStdin(fileFlag, regExp)
				utilities.HandleFile(os.Args[1], stdinMap, regExp)
			}
		} else {
			usage.PrintUsage()
		}

	} else if (stat.Mode() & os.ModeCharDevice) == 0 {
		fileFlag = false
		utilities.HandleStdin(fileFlag, regExp)

	} else if len(os.Args) == 4 {
		if os.Args[1] == "-e" {
			regExp = os.Args[2]
			fileFlag = true
			utilities.HandleFile(os.Args[3], nil, regExp)
		} else if strings.Contains(os.Args[1], "-") {
			fmt.Println("Unknown flag provided...")
			usage.PrintUsage()
		} else if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
			fmt.Println("Error: File does not exist")
			return
		} else {
			utilities.HandleFile(os.Args[1], nil, regExp)
		}

	} else if len(os.Args) == 3 {
		if os.Args[1] == "-e" {
			fmt.Println("Expression and file is not provided...")
			usage.PrintUsage()
		}
	} else if len(os.Args) == 2 {
		if os.Args[1] == "-h" || os.Args[1] == "help" {
			usage.PrintUsage()
		} else if os.Args[1] == "-e" {
			fmt.Println("No expression is provided...")
			usage.PrintUsage()
		} else if strings.Contains(os.Args[1], "-") {
			fmt.Println("Unknown flag provided...")
			usage.PrintUsage()
		} else if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
			fmt.Println("Error: File does not exist")
			return
		} else {
			utilities.HandleFile(os.Args[1], nil, regExp)
		}
	} else {
		usage.PrintUsage()
	}
}

package usage

import "fmt"

// For printing usage of the seek
func PrintUsage() {
	fmt.Print("\nA tool to process and write unique lines written by Aakansh\n\n")
	fmt.Print("Usage: seek <filename>\n\n")
	fmt.Print("Basic Functioning: \n\n")
	fmt.Println("  seek file.txt						(Read and write to the same file)")
	fmt.Println("  cat file.txt | seek					(Read from stdin and display to the stdout)")
	fmt.Println("  cat file.txt | seek newfile.txt			(Read from stdin and write to a specific file)")
	fmt.Print("\nAdvanced Functioning: \n\n")
	fmt.Println("  seek -e 'expression'					(Regular expressions for filtering)")
}

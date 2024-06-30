package runner

import (
	"fmt"
	"os"
)

// For printing usage of the refine
func PrintUsage() {
	h := "\nUsage: refine [options]\n\n"
	h += "Options: [flag] [argument] [Description]\n\n"
	h += "DIRECT:\n"
	h += "  refine file.txt\t\t\t(Read and write to the same file)\n"
	h += "  refine file1.txt file2.txt\t\t(Read from file1 and write to the file2)\n\n"
	h += "STDIN:\n"
	h += "  cat file.txt | refine\t\t\t(Read from stdin and display to stdout)\n"
	h += "  cat file.txt | refine newfile.txt\t(Read from stdin and write to a specific file)\n\n"
	h += "FEATURES: (ONLY DIRECT MODE)\n"
	h += "  refine -w, --wildcard\t\t\t(Sort all files in the directory)\n"
	h += "  refine -we, --wildcard-exception\t(Specify files to be skipped while using wildcard)\n\n"
	h += "DEBUG:\n"
	h += "  refine -v, --version\t\t\t(Check current version)"

	fmt.Fprintln(os.Stderr, h)
}

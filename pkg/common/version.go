package common

import (
	"fmt"
	"os"
)

// Function for checking the version info
func CheckVersion() {
	fmt.Fprintln(os.Stderr, Succfix, "Current refine version:", G.BrGreen("v1.0"))
}

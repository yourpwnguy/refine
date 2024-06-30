package runner

import (
	"fmt"
	"os"
)

// Function for checking the version info
func CheckVersion() {
	fmt.Fprintln(os.Stderr, Succfix, "Current refine version:", g.BrGreen("v1.0"))
}

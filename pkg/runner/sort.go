package runner

import (
	"slices"
)

// Function for sorting all the lines
func SortMap(temp *map[string]struct{}) []string {
	var sorted []string
	for line := range *temp {
		sorted = append(sorted, line)
	}

	// Sort the slice in-place
	slices.Sort(sorted)
	return sorted
}

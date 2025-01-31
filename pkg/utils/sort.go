package utils

import (
	"sort"
)

// Function for sorting all the lines
func SortMap(temp map[string]struct{}) []string {
	sorted := make([]string, 0, len(temp))
	for line := range temp {
		sorted = append(sorted, line)
	}

	// Sort the slice in-place
	sort.Strings(sorted)
	return sorted
}

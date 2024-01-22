package utilities

import (
	"regexp"
	"strings"
)

// For processing the regular expressions -e
func ExtractUsingRegex(temp map[string]struct{}, regExp string) map[string]struct{} {
	filtered := make(map[string]struct{})
	re := regexp.MustCompile(regExp)

	for key := range temp {
		if strings.TrimSpace(key) != "" {
			matches := re.FindAllStringSubmatch(key, -1)
			for _, match := range matches {
				if len(match) >= 1 {
					extractedPart := match[0]
					filtered[extractedPart] = struct{}{}
				}
			}
		}
	}
	return filtered
}
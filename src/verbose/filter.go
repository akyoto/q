package verbose

import "strings"

// filter returns true if the function doesn't pass the filter.
func filter(fullName string, pattern string) bool {
	if pattern == "" {
		return false
	}

	if strings.Contains(pattern, ".") {
		return !strings.HasPrefix(fullName, pattern)
	}

	return !strings.HasPrefix(fullName, "main."+pattern)
}
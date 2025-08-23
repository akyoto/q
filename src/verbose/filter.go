package verbose

import "strings"

// filter returns true if the function doesn't pass the filter.
func filter(fullName string, filter string) bool {
	if filter == "" {
		return false
	}

	if strings.Contains(filter, ".") {
		return fullName != filter
	}

	return !strings.Contains(fullName, filter)
}
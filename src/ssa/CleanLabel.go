package ssa

import "strings"

// CleanLabel removes the function name from the label.
func CleanLabel(label string) string {
	pos := strings.IndexByte(label, ':')

	if pos != -1 && pos != len(label)-1 {
		return label[pos+1:]
	}

	return label
}
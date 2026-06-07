package cli

import (
	_ "embed"
)

//go:embed keywords.txt
var keywordsText string

// keywords shows the entire list of keywords.
func keywords() int {
	show(keywordsText)
	return success
}
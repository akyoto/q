package cli

import (
	_ "embed"
)

//go:embed help.txt
var helpText string

// help shows the command line argument usage.
func help() int {
	show(helpText)
	return 0
}
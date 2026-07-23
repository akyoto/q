package cli

import (
	_ "embed"
)

//go:embed help.txt
var helpText string

// help shows the command line argument usage.
func help() int {
	render(helpText)
	return success
}
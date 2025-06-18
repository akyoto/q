package cli

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed help.txt
var helpText string

// help shows the command line argument usage.
func help() int {
	fmt.Fprintln(os.Stdout, helpText)
	return 0
}
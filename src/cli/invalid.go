package cli

import (
	"fmt"
	"os"
)

// invalid shows the help and returns exit code 2 (invalid parameters).
func invalid() int {
	fmt.Fprintln(os.Stderr, helpText)
	return 2
}
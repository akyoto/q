package cli

import (
	"os"

	"git.urbach.dev/go/color"
)

// invalid shows the help on stderr and returns exit code 2 (invalid parameters).
func invalid() int {
	color.Redirect(os.Stderr)
	help()
	return 2
}
package cli

import (
	_ "embed"
)

//go:embed version.txt
var versionText string

// version shows the commit date and hash.
func version() int {
	show(versionText)
	return 0
}
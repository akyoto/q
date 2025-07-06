//go:build !linux

package memfile

import (
	"os"

	"git.urbach.dev/cli/q/src/global"
)

// New creates a new anonymous in-memory file.
func New(name string) (*os.File, error) {
	pattern := ""

	if global.OS == "windows" {
		pattern = "*.exe"
	}

	return os.CreateTemp("", pattern)
}
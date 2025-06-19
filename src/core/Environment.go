package core

import (
	"git.urbach.dev/cli/q/src/fs"
)

// Environment holds information about the entire build.
type Environment struct {
	Functions map[string]*Function
	Files     []*fs.File
}
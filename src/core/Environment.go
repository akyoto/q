package core

import (
	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/fs"
)

// Environment holds information about the entire build.
type Environment struct {
	Build     *build.Build
	Functions map[string]*Function
	Files     []*fs.File
}
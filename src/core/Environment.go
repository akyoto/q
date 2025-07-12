package core

import (
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/fs"
)

// Environment holds information about the entire build.
// Functions can access information about other functions using this.
// We'll also pass this to the linker because it contains everything
// that's needed to generate an executable file.
type Environment struct {
	Build     *config.Build
	Init      *Function
	Main      *Function
	Functions map[string]*Function
	Files     []*fs.File
}
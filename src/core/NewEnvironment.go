package core

import (
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/fs"
)

// NewEnvironment creates a new environment.
func NewEnvironment(build *config.Build) *Environment {
	return &Environment{
		Build:    build,
		Files:    make([]*fs.File, 0, 16),
		Packages: make(map[string]*Package, 8),
	}
}
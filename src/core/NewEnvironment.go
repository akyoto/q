package core

import (
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/types"
)

// NewEnvironment creates a new environment.
func NewEnvironment(build *config.Build) *Environment {
	env := &Environment{
		Build:    build,
		Files:    make([]*fs.File, 0, 16),
		Packages: make(map[string]*Package, 8),
	}

	env.pointerTypes.Store(types.Any, types.AnyPointer)
	env.pointerTypes.Store(types.Byte, types.CString)
	env.sliceTypes.Store(types.Byte, types.String)
	return env
}
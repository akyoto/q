package core

import (
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/types"
)

// NewEnvironment creates a new environment.
func NewEnvironment(build *config.Build) *Environment {
	return &Environment{
		Build:    build,
		Files:    make([]*fs.File, 0, 16),
		Packages: make(map[string]*Package, 8),
		PointerTypes: map[types.Type]types.Type{
			types.Any:  types.AnyPointer,
			types.Byte: types.CString,
		},
		ResourceTypes: map[types.Type]types.Type{},
		SliceTypes: map[types.Type]types.Type{
			types.Byte: types.String,
		},
	}
}
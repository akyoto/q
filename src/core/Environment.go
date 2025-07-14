package core

import (
	"iter"

	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/fs"
)

// Environment holds information about the entire build.
// Functions can access information about other functions using this.
// We'll also pass this to the linker because it contains everything
// that's needed to generate an executable file.
type Environment struct {
	Build        *config.Build
	Init         *Function
	Main         *Function
	Packages     map[string]*Package
	Files        []*fs.File
	NumFunctions int
}

// Function looks up a function by the package name and raw function name.
func (env *Environment) Function(pkgName string, name string) *Function {
	pkg, exists := env.Packages[pkgName]

	if !exists {
		return nil
	}

	fn, exists := pkg.Functions[name]

	if !exists {
		return nil
	}

	return fn
}

// Functions returns an iterator over all functions.
func (env *Environment) Functions() iter.Seq[*Function] {
	return func(yield func(*Function) bool) {
		for _, pkg := range env.Packages {
			for _, fn := range pkg.Functions {
				if !yield(fn) {
					return
				}
			}
		}
	}
}
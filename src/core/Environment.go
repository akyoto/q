package core

import (
	"iter"

	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/types"
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
	TypeCache
}

// AddPackage returns an existing package with the giving name or creates a new one.
func (env *Environment) AddPackage(name string, isExtern bool) *Package {
	pkg, exists := env.Packages[name]

	if !exists {
		pkg = &Package{
			Name:      name,
			Constants: make(map[string]*Constant),
			Functions: make(map[string]*Function, 8),
			Structs:   make(map[string]*types.Struct),
			Globals:   make(map[string]*Global),
			IsExtern:  isExtern,
		}

		env.Packages[name] = pkg
	}

	return pkg
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
				for variant := range fn.Variants {
					if !yield(variant) {
						return
					}
				}
			}
		}
	}
}

// Globals returns an iterator over all globals.
func (env *Environment) Globals() iter.Seq[*Global] {
	return func(yield func(*Global) bool) {
		for _, pkg := range env.Packages {
			for _, global := range pkg.Globals {
				if !yield(global) {
					return
				}
			}
		}
	}
}

// LiveFunctions returns an iterator over functions that are alive,
// starting with `run.init` and all of its dependencies.
func (env *Environment) LiveFunctions() iter.Seq[*Function] {
	return func(yield func(*Function) bool) {
		running := true
		traversed := make(map[*Function]bool, env.NumFunctions)

		env.Init.EachDependency(traversed, func(f *Function) {
			if !running {
				return
			}

			running = yield(f)
		})
	}
}

// ResolveTypes resolves all the type tokens in structs, globals and function parameters.
func (env *Environment) ResolveTypes() error {
	err := env.parseStructs(env.Structs())

	if err != nil {
		return err
	}

	err = env.parseGlobals(env.Globals())

	if err != nil {
		return err
	}

	return env.parseParameters(env.Functions())
}

// Structs returns an iterator over all structs.
func (env *Environment) Structs() iter.Seq[*types.Struct] {
	return func(yield func(*types.Struct) bool) {
		for _, pkg := range env.Packages {
			for _, structure := range pkg.Structs {
				if !yield(structure) {
					return
				}
			}
		}
	}
}
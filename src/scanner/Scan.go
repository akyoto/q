package scanner

import (
	"path/filepath"

	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/global"
	"git.urbach.dev/cli/q/src/types"
)

// Scan scans all the files included in the build.
func Scan(build *config.Build) (*core.Environment, error) {
	s := scanner{
		constants: make(chan *core.Constant, 8),
		functions: make(chan *core.Function, 8),
		files:     make(chan *fs.File, 8),
		structs:   make(chan *types.Struct, 8),
		globals:   make(chan *core.Global, 1),
		errors:    make(chan error),
		build:     build,
	}

	go func() {
		s.queueDirectory(filepath.Join(global.Library, "run"), "run")
		s.queue(build.Files...)
		s.group.Wait()
		close(s.constants)
		close(s.functions)
		close(s.files)
		close(s.structs)
		close(s.globals)
		close(s.errors)
	}()

	env := core.NewEnvironment(build)

	for s.functions != nil || s.files != nil || s.structs != nil || s.constants != nil || s.globals != nil || s.errors != nil {
		select {
		case f, ok := <-s.functions:
			if !ok {
				s.functions = nil
				continue
			}

			env.ReceiveFunction(f)

		case file, ok := <-s.files:
			if !ok {
				s.files = nil
				continue
			}

			env.ReceiveFile(file)

		case structure, ok := <-s.structs:
			if !ok {
				s.structs = nil
				continue
			}

			env.ReceiveStruct(structure)

		case constant, ok := <-s.constants:
			if !ok {
				s.constants = nil
				continue
			}

			env.ReceiveConstant(constant)

		case global, ok := <-s.globals:
			if !ok {
				s.globals = nil
				continue
			}

			env.ReceiveGlobal(global)

		case err, ok := <-s.errors:
			if !ok {
				s.errors = nil
				continue
			}

			return env, err
		}
	}

	return env, nil
}
package scanner

import (
	"path/filepath"

	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/global"
)

// Scan scans all the files included in the build.
func Scan(build *config.Build) (*core.Environment, error) {
	s := scanner{
		constants: make(chan *core.Constant),
		functions: make(chan *core.Function),
		files:     make(chan *fs.File),
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
		close(s.errors)
	}()

	all := &core.Environment{
		Build:    build,
		Files:    make([]*fs.File, 0, 8),
		Packages: make(map[string]*core.Package, 8),
	}

	for s.functions != nil || s.files != nil || s.constants != nil || s.errors != nil {
		select {
		case f, ok := <-s.functions:
			if !ok {
				s.functions = nil
				continue
			}

			f.All = all
			pkg := all.AddPackage(f.Package, f.IsExtern())
			pkg.Functions[f.Name] = f
			all.NumFunctions++

		case file, ok := <-s.files:
			if !ok {
				s.files = nil
				continue
			}

			all.Files = append(all.Files, file)

		case constant, ok := <-s.constants:
			if !ok {
				s.constants = nil
				continue
			}

			pkg := all.AddPackage(constant.File.Package, false)
			pkg.Constants[constant.Name] = constant

		case err, ok := <-s.errors:
			if !ok {
				s.errors = nil
				continue
			}

			return all, err
		}
	}

	return all, nil
}
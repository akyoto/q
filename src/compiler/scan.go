package compiler

import (
	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/scanner"
)

func scan(b *build.Build) (*core.Environment, error) {
	functions, files, errs := scanner.Scan(b)

	all := &core.Environment{
		Files:     make([]*fs.File, 0, 8),
		Functions: make(map[string]*core.Function, 32),
	}

	for functions != nil || files != nil || errs != nil {
		select {
		case f, ok := <-functions:
			if !ok {
				functions = nil
				continue
			}

			all.Functions[f.String()] = f

		case file, ok := <-files:
			if !ok {
				files = nil
				continue
			}

			all.Files = append(all.Files, file)

		case err, ok := <-errs:
			if !ok {
				errs = nil
				continue
			}

			return all, err
		}
	}

	return all, nil
}
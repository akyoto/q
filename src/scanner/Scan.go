package scanner

import (
	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/fs"
)

// Scan scans all the files included in the build.
func Scan(b *build.Build) (*core.Environment, error) {
	s := scanner{
		functions: make(chan *core.Function),
		files:     make(chan *fs.File),
		errors:    make(chan error),
		build:     b,
	}

	go func() {
		s.queue(b.Files...)
		s.group.Wait()
		close(s.functions)
		close(s.files)
		close(s.errors)
	}()

	all := &core.Environment{
		Files:     make([]*fs.File, 0, 8),
		Functions: make(map[string]*core.Function, 32),
	}

	for s.functions != nil || s.files != nil || s.errors != nil {
		select {
		case f, ok := <-s.functions:
			if !ok {
				s.functions = nil
				continue
			}

			all.Functions[f.String()] = f

		case file, ok := <-s.files:
			if !ok {
				s.files = nil
				continue
			}

			all.Files = append(all.Files, file)

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
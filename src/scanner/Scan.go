package scanner

import (
	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/fs"
)

// Scan scans all the files included in the build.
func Scan(b *build.Build) (<-chan *core.Function, <-chan *fs.File, <-chan error) {
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

	return s.functions, s.files, s.errors
}
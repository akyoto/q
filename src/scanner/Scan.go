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
		items: make(chan any, 512),
		build: build,
	}

	go func() {
		s.queueDirectory(filepath.Join(global.Library, "run"), "run")
		s.queueDirectory(filepath.Join(global.Library, "strings"), "strings")
		s.queue(build.Files...)
		s.group.Wait()
		close(s.items)
	}()

	env := core.NewEnvironment(build)

	for item := range s.items {
		switch v := item.(type) {
		case *core.Function:
			env.ReceiveFunction(v)
		case *fs.File:
			env.ReceiveFile(v)
		case *types.Struct:
			env.ReceiveStruct(v)
		case *core.Constant:
			env.ReceiveConstant(v)
		case *types.Enum:
			env.ReceiveEnum(v)
		case *core.Global:
			env.ReceiveGlobal(v)
		case error:
			return env, v
		}
	}

	return env, nil
}
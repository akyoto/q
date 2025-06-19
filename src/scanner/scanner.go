package scanner

import (
	"sync"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/fs"
)

// scanner is used to scan files before the actual compilation step.
type scanner struct {
	functions chan *core.Function
	files     chan *fs.File
	errors    chan error
	build     *build.Build
	queued    sync.Map
	group     sync.WaitGroup
}
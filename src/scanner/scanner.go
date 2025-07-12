package scanner

import (
	"sync"

	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/fs"
)

// scanner is used to scan files before the actual compilation step.
type scanner struct {
	functions chan *core.Function
	files     chan *fs.File
	errors    chan error
	build     *config.Build
	queued    sync.Map
	group     sync.WaitGroup
}
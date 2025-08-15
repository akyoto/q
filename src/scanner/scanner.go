package scanner

import (
	"sync"

	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/types"
)

// scanner is used to scan files before the actual compilation step.
type scanner struct {
	constants chan *core.Constant
	functions chan *core.Function
	files     chan *fs.File
	structs   chan *types.Struct
	errors    chan error
	build     *config.Build
	queued    sync.Map
	group     sync.WaitGroup
}
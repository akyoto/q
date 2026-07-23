package scanner

import (
	"sync"

	"git.urbach.dev/cli/q/src/config"
)

// scanner is used to scan files before the actual compilation step.
type scanner struct {
	items  chan any
	build  *config.Build
	queued sync.Map
	group  sync.WaitGroup
}
package fs

import (
	"sync/atomic"

	"git.urbach.dev/cli/q/src/token"
)

// Import contains data about an import statement in a file.
type Import struct {
	Package  string
	Position token.Position
	Used     atomic.Uint64
}

// IsUsed returns true if the import was used.
func (imp *Import) IsUsed() bool {
	return imp.Used.Load() > 0
}
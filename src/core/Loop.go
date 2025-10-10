package core

import (
	"git.urbach.dev/cli/q/src/ssa"
)

// Loop is a loop on the stack that remembers all active loops.
type Loop struct {
	Head         *ssa.Block
	Exit         *ssa.Block
	FromValue    ssa.Value
	IteratorName string
}
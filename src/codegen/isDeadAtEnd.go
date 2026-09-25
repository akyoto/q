package codegen

import (
	"git.urbach.dev/cli/q/src/ssa"
)

// isDeadAtEnd returns true if every successor of `pre` is `live.Block`.
func (f *Function) isDeadAtEnd(live *Step, pre *ssa.Block) bool {
	succs := f.BlockToSuccessors[pre]

	if len(succs) == 0 {
		return false
	}

	for _, succ := range succs {
		if succ != live.Block {
			return false
		}
	}

	return true
}
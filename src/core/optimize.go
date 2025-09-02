package core

import (
	"git.urbach.dev/cli/q/src/fold"
	"git.urbach.dev/cli/q/src/ssa"
)

// optimize applies various algorithms after the compilation.
func (f *Function) optimize() error {
	// Sometimes unreachable blocks are created at the end of
	// an infinite loop. These should be removed to avoid
	// unnecessary return statements in the later phases.
	f.removeDeadBlocks()

	// Binary operations with constant operands are evaluated
	// at compile time. For example, 1 + 2 becomes 3, and the
	// result is propagated to subsequent operations.
	var folded map[ssa.Value]struct{}

	if f.Env.Build.FoldConstants {
		folded = fold.Constants(f.IR)
	}

	// After cleaning up some of the instructions we can proceed
	// to calculate the list of users.
	f.ComputeUsers()

	// Now that we have the list of users for each instruction,
	// we can filter out dead values.
	err := f.removeDeadCode(folded)

	if err != nil {
		return err
	}

	// Resource types that are still defined at the end of a
	// scope must be freed.
	return f.checkResources()
}
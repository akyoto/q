package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fold"
	"git.urbach.dev/cli/q/src/ssa"
)

// optimize applies various algorithms after the compilation.
func (f *Function) optimize() error {
	// Sometimes unreachable blocks are created at the end of
	// an infinite loop. These should be removed to avoid
	// unnecessary return statements in the later phases.
	f.removeDeadBlocks()

	// After the removal of dead blocks, if the last block is
	// not part of a loop and did not end with a return
	// statement, an implicit return is inserted.
	if f.needsReturn() {
		if len(f.Output) > 0 {
			return errors.New(&ReturnCountMismatch{Count: 0, ExpectedCount: len(f.Output)}, f.File, f.Output[0].StartPos)
		}

		f.Block().Append(&ssa.Return{})
	}

	// Copies were inserted for assignments to be safe in case
	// loops would replace an existing value with a phi inside
	// the loop. Now that all the loop replacements happened,
	// we can safely remove the copies.
	if f.Env.Build.RemoveCopies {
		f.removeCopies()
	}

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
	return f.verifyDeallocation()
}
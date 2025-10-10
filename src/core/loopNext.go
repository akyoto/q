package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// loopNext starts the next iteration of the loop.
func (f *Function) loopNext(loop *Loop) {
	if loop.FromValue != nil {
		one := f.Append(&ssa.Int{Int: 1})

		nextIteration := f.Append(&ssa.BinaryOp{
			Op:    token.Add,
			Left:  loop.FromValue,
			Right: one,
		})

		f.Block().Identify(loop.IteratorName, nextIteration)
	}

	f.jump(loop.Head)
}
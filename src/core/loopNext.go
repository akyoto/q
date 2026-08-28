package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// loopNext starts the next iteration of the loop.
func (f *Function) loopNext(loop *Loop) {
	if loop.IteratorName != "" {
		endOfLoopValue, _, _ := f.findIdentifier(loop.IteratorName)
		one := f.Append(&ssa.Int{Int: 1})

		nextIteration := &ssa.BinaryOp{
			Op:    token.Add,
			Left:  endOfLoopValue,
			Right: one,
		}

		f.Block().Append(nextIteration)
		f.Block().Identify(loop.IteratorName, nextIteration)
	}

	f.jump(loop.Head)
}
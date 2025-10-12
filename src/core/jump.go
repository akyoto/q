package core

import "git.urbach.dev/cli/q/src/ssa"

// jump inserts a `jump` instruction to the block and adds it as a successor
// if the last instruction was not a `jump` or `return`.
func (f *Function) jump(block *ssa.Block) {
	switch f.Block().Last().(type) {
	case *ssa.Jump, *ssa.Return:
		return
	}

	f.Block().Append(&ssa.Jump{To: block})
	f.Block().AddSuccessor(block)
}
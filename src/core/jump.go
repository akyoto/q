package core

import "git.urbach.dev/cli/q/src/ssa"

// jump inserts a `jump` instruction to the block and adds it as a successor.
// It performs a check to see if the last instruction was a `return` instruction.
// In case `return` was found, it will not do anything.
func (f *Function) jump(block *ssa.Block) {
	_, returned := f.Block().Last().(*ssa.Return)

	if returned {
		return
	}

	f.Block().Append(&ssa.Jump{To: block})
	f.Block().AddSuccessor(block)
}
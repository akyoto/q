package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// Loop compiles an endless loop.
func (f *Function) Loop(tokens token.List) error {
	blockStart, blockEnd, err := f.block(tokens)

	if err != nil {
		return err
	}

	f.Count.Loop++
	bodyLabel := f.CreateLabel("loop body", f.Count.Loop)
	exitLabel := f.CreateLabel("loop exit", f.Count.Loop)
	loopBody := ssa.NewBlock(bodyLabel)
	loopExit := ssa.NewBlock(exitLabel)
	f.AddBlock(loopBody)
	err = f.compileTokens(tokens[blockStart+1 : blockEnd])

	if err != nil {
		return err
	}

	f.Append(&ssa.Jump{To: loopBody})
	f.AddBlock(loopExit)
	return nil
}
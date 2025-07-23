package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// If compiles a branch instruction.
func (f *Function) If(tokens token.List) error {
	blockStart, blockEnd, err := f.block(tokens)

	if err != nil {
		return err
	}

	condition := expression.Parse(tokens[1:blockStart])

	f.Count.Branch++
	thenLabel := f.CreateLabel("if.then", f.Count.Branch)
	elseLabel := f.CreateLabel("if.else", f.Count.Branch)
	thenBlock := ssa.NewBlock(thenLabel)
	elseBlock := ssa.NewBlock(elseLabel)
	f.Block().AddSuccessor(thenBlock)
	f.Block().AddSuccessor(elseBlock)
	err = f.compileCondition(condition, thenBlock, elseBlock)

	if err != nil {
		return err
	}

	f.AddBlock(thenBlock)
	err = f.compileTokens(tokens[blockStart+1 : blockEnd])

	if err != nil {
		return err
	}

	f.AddBlock(elseBlock)
	return nil
}
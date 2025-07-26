package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/ssa"
)

// compileIf compiles a branch instruction.
func (f *Function) compileIf(branch *ast.If) error {
	f.Count.Branch++
	thenLabel := f.CreateLabel("if.then", f.Count.Branch)
	elseLabel := f.CreateLabel("if.else", f.Count.Branch)
	thenBlock := ssa.NewBlock(thenLabel)
	elseBlock := ssa.NewBlock(elseLabel)
	f.Block().AddSuccessor(thenBlock)
	f.Block().AddSuccessor(elseBlock)
	thenBlock.AddSuccessor(elseBlock)
	err := f.compileCondition(branch.Condition, thenBlock, elseBlock)

	if err != nil {
		return err
	}

	f.AddBlock(thenBlock)
	err = f.compileAST(branch.Body)

	if err != nil {
		return err
	}

	f.AddBlock(elseBlock)
	return nil
}
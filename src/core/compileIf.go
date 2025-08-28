package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/ssa"
)

// compileIf compiles a branch instruction.
func (f *Function) compileIf(branch *ast.If) error {
	f.Count.Branch++
	thenLabel := f.CreateLabel("if.then", f.Count.Branch)
	exitLabel := f.CreateLabel("if.exit", f.Count.Branch)
	thenBlock := ssa.NewBlock(thenLabel)
	exitBlock := ssa.NewBlock(exitLabel)
	beforeIf := f.Block()
	beforeIf.AddSuccessor(thenBlock)

	if branch.Else == nil {
		beforeIf.AddSuccessor(exitBlock)
		err := f.compileCondition(branch.Condition, thenBlock, exitBlock)

		if err != nil {
			return err
		}

		// Append the if.then block
		f.AddBlock(thenBlock)
		err = f.compileAST(branch.Body)

		if err != nil {
			return err
		}

		f.jump(exitBlock)
	} else {
		elseLabel := f.CreateLabel("if.else", f.Count.Branch)
		elseBlock := ssa.NewBlock(elseLabel)
		beforeIf.AddSuccessor(elseBlock)
		err := f.compileCondition(branch.Condition, thenBlock, elseBlock)

		if err != nil {
			return err
		}

		// Append the if.then block
		f.AddBlock(thenBlock)
		err = f.compileAST(branch.Body)

		if err != nil {
			return err
		}

		f.jump(exitBlock)

		// Append the if.else block
		f.AddBlock(elseBlock)
		err = f.compileAST(branch.Else)

		if err != nil {
			return err
		}

		f.jump(exitBlock)
	}

	f.AddBlock(exitBlock)
	return nil
}
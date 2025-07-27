package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/ssa"
)

// compileIf compiles a branch instruction.
func (f *Function) compileIf(branch *ast.If) error {
	f.Count.Branch++
	thenLabel := f.CreateLabel("if.then", f.Count.Branch)
	endLabel := f.CreateLabel("if.end", f.Count.Branch)
	thenBlock := ssa.NewBlock(thenLabel)
	endBlock := ssa.NewBlock(endLabel)
	beforeIf := f.Block()
	beforeIf.AddSuccessor(thenBlock)
	thenBlock.AddSuccessor(endBlock)

	if branch.Else == nil {
		err := f.compileCondition(branch.Condition, thenBlock, endBlock)

		if err != nil {
			return err
		}

		// Append the if.then block
		f.AddBlock(thenBlock)
		err = f.compileAST(branch.Body)

		if err != nil {
			return err
		}

		beforeIf.AddSuccessor(endBlock)
	} else {
		elseLabel := f.CreateLabel("if.else", f.Count.Branch)
		elseBlock := ssa.NewBlock(elseLabel)
		err := f.compileCondition(branch.Condition, thenBlock, elseBlock)

		if err != nil {
			return err
		}

		beforeIf.AddSuccessor(elseBlock)
		elseBlock.AddSuccessor(endBlock)

		// Append the if.then block
		f.AddBlock(thenBlock)
		err = f.compileAST(branch.Body)

		if err != nil {
			return err
		}

		f.Append(&ssa.Jump{To: endBlock})

		// Append the if.else block
		f.AddBlock(elseBlock)
		err = f.compileAST(branch.Else)

		if err != nil {
			return err
		}
	}

	f.AddBlock(endBlock)
	return nil
}
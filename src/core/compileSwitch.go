package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/ssa"
)

// compileSwitch compiles a multi-branch instruction.
func (f *Function) compileSwitch(s *ast.Switch) error {
	f.Count.Switch++
	exitLabel := f.CreateLabel("switch.exit", f.Count.Switch)
	exitBlock := ssa.NewBlock(exitLabel)

	for i, branch := range s.Cases {
		if branch.Condition == nil {
			err := f.compileAST(branch.Body)

			if err != nil {
				return err
			}

			f.jump(exitBlock)
			f.AddBlock(exitBlock)
			break
		}

		f.Count.Branch++
		thenLabel := f.CreateLabel("if.then", f.Count.Branch)
		thenBlock := ssa.NewBlock(thenLabel)
		var elseBlock *ssa.Block

		if i < len(s.Cases)-1 {
			elseLabel := f.CreateLabel("if.else", f.Count.Branch)
			elseBlock = ssa.NewBlock(elseLabel)
		} else {
			elseBlock = exitBlock
		}

		caseBlock := f.Block()
		caseBlock.AddSuccessor(thenBlock)
		caseBlock.AddSuccessor(elseBlock)
		err := f.compileCondition(branch.Condition, thenBlock, elseBlock)

		if err != nil {
			return err
		}

		f.AddBlock(thenBlock)
		err = f.compileAST(branch.Body)

		if err != nil {
			return err
		}

		f.jump(exitBlock)
		f.AddBlock(elseBlock)
	}

	return nil
}
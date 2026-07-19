package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/ssa"
)

// compileSwitch compiles a multi-branch instruction.
func (f *Function) compileSwitch(s *ast.Switch) error {
	f.Count.Switch++

	var (
		head      ssa.Value
		err       error
		exitLabel = f.CreateLabel("switch.exit", f.Count.Switch)
		exitBlock = ssa.NewBlock(exitLabel)
	)

	if s.Head != nil {
		head, err = f.evaluateRight(s.Head)

		if err != nil {
			return err
		}
	}

	for i, branch := range s.Cases {
		if branch.Condition == nil {
			before := f.Block().Identifiers.Before
			err = f.compileAST(branch.Body)

			if err != nil {
				return err
			}

			f.deleteResources(before)
			f.jump(exitBlock)
			f.AddBlock(exitBlock)
			break
		}

		f.Count.Branch++
		thenLabel := f.CreateLabel("case.then", f.Count.Branch)
		thenBlock := ssa.NewBlock(thenLabel)
		var elseBlock *ssa.Block

		if i < len(s.Cases)-1 {
			elseLabel := f.CreateLabel("case.else", f.Count.Branch)
			elseBlock = ssa.NewBlock(elseLabel)
		} else {
			elseBlock = exitBlock
		}

		if head != nil {
			caseValue, err := f.evaluateRight(branch.Condition)

			if err != nil {
				return err
			}

			condition, err := f.equal(head, caseValue, branch.Condition.Source())

			if err != nil {
				return err
			}

			block := f.Block()
			block.AddSuccessor(thenBlock)
			block.AddSuccessor(elseBlock)

			block.Append(&ssa.Branch{
				Condition: condition,
				Then:      thenBlock,
				Else:      elseBlock,
			})
		} else {
			err = f.compileCondition(branch.Condition, thenBlock, elseBlock)

			if err != nil {
				return err
			}
		}

		f.AddBlock(thenBlock)
		before := f.Block().Identifiers.Before
		err = f.compileAST(branch.Body)

		if err != nil {
			return err
		}

		f.deleteResources(before)
		f.jump(exitBlock)
		f.AddBlock(elseBlock)
	}

	return nil
}
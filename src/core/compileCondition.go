package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// compileCondition inserts code to jump to the start label or end label depending on the truth of the condition.
func (f *Function) compileCondition(condition *expression.Expression, thenBlock *ssa.Block, elseBlock *ssa.Block) error {
	switch condition.Token.Kind {
	case token.LogicalOr:
		f.Count.SubBranch++
		leftFailLabel := f.CreateLabel("or", f.Count.SubBranch)
		leftFail := ssa.NewBlock(leftFailLabel)
		f.Block().AddSuccessor(leftFail)

		// Left
		left := condition.Children[0]
		err := f.compileCondition(left, thenBlock, leftFail)

		if err != nil {
			return err
		}

		f.AddBlock(leftFail)

		// Right
		right := condition.Children[1]
		err = f.compileCondition(right, thenBlock, elseBlock)
		return err

	case token.LogicalAnd:
		f.Count.SubBranch++
		leftSuccessLabel := f.CreateLabel("and", f.Count.SubBranch)
		leftSuccess := ssa.NewBlock(leftSuccessLabel)
		f.Block().AddSuccessor(leftSuccess)

		// Left
		left := condition.Children[0]
		err := f.compileCondition(left, leftSuccess, elseBlock)

		if err != nil {
			return err
		}

		f.AddBlock(leftSuccess)

		// Right
		right := condition.Children[1]
		err = f.compileCondition(right, thenBlock, elseBlock)

		return err

	case token.Equal, token.NotEqual, token.Greater, token.Less, token.GreaterEqual, token.LessEqual:
		conditionValue, err := f.eval(condition)

		if err != nil {
			return err
		}

		branch := &ssa.Branch{
			Condition: conditionValue,
			Then:      thenBlock,
			Else:      elseBlock,
		}

		f.Append(branch)
		return nil

	default:
		return errors.New(InvalidCondition, f.File, condition.Token.Position)
	}
}
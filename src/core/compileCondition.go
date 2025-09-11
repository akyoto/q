package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// compileCondition inserts code to jump to the start label or end label depending on the truth of the condition.
func (f *Function) compileCondition(condition *expression.Expression, thenBlock *ssa.Block, elseBlock *ssa.Block) error {
	switch condition.Token.Kind {
	case token.LogicalOr:
		f.Count.SubBranch++
		leftFailLabel := f.CreateLabel("or", f.Count.SubBranch)
		leftFail := ssa.NewBlock(leftFailLabel)

		// Left
		left := condition.Children[0]
		err := f.compileCondition(left, thenBlock, leftFail)

		if err != nil {
			return err
		}

		f.Block().AddSuccessor(leftFail)
		f.AddBlock(leftFail)

		// Right
		right := condition.Children[1]
		err = f.compileCondition(right, thenBlock, elseBlock)
		return err

	case token.LogicalAnd:
		f.Count.SubBranch++
		leftSuccessLabel := f.CreateLabel("and", f.Count.SubBranch)
		leftSuccess := ssa.NewBlock(leftSuccessLabel)

		// Left
		left := condition.Children[0]
		err := f.compileCondition(left, leftSuccess, elseBlock)

		if err != nil {
			return err
		}

		f.Block().AddSuccessor(leftSuccess)
		f.AddBlock(leftSuccess)

		// Right
		right := condition.Children[1]
		err = f.compileCondition(right, thenBlock, elseBlock)

		return err

	case token.Equal, token.NotEqual, token.Greater, token.Less, token.GreaterEqual, token.LessEqual:
		conditionValue, err := f.evaluate(condition)

		if err != nil {
			return err
		}

		comparison := conditionValue.(*ssa.BinaryOp)
		left := comparison.Left

		if left.Type() == types.Error {
			right := comparison.Right.(*ssa.Int)

			switch {
			case condition.Token.Kind == token.NotEqual && right.Int == 0:
				for _, protected := range thenBlock.Protected[left] {
					thenBlock.Unidentify(protected)
				}

				thenBlock.Unprotect(left)
				elseBlock.Unidentify(left)
				elseBlock.Unprotect(left)
			case condition.Token.Kind == token.Equal && right.Int == 0:
				for _, protected := range elseBlock.Protected[left] {
					elseBlock.Unidentify(protected)
				}

				elseBlock.Unprotect(left)
				thenBlock.Unidentify(left)
				thenBlock.Unprotect(left)
			}
		}

		branch := &ssa.Branch{
			Condition: conditionValue,
			Then:      thenBlock,
			Else:      elseBlock,
		}

		f.Append(branch)
		return nil

	case token.Not:
		return f.compileCondition(condition.Children[0], elseBlock, thenBlock)

	default:
		if condition.Token.Kind.IsAssignment() {
			return errors.New(InvalidCondition, f.File, condition.Token.Position)
		}

		value, err := f.evaluate(condition)

		if err != nil {
			return err
		}

		if !types.Is(value.Type(), types.Bool) {
			return errors.New(InvalidCondition, f.File, condition.Source().StartPos)
		}

		branch := &ssa.Branch{
			Condition: value,
			Then:      thenBlock,
			Else:      elseBlock,
		}

		f.Append(branch)
		return nil
	}
}
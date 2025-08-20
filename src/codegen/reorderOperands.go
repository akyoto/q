package codegen

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// reorderOperands reorders the left and right operands of a binary operation
// so that the left operand matches the destination register. It will also
// try to match the right operand with an integer immediate.
func (f *Function) reorderOperands(step *Step, instr *ssa.BinaryOp) {
	if instr.Op != token.Add {
		return
	}

	leftStep := f.ValueToStep[instr.Left]
	rightStep := f.ValueToStep[instr.Right]

	if leftStep.Register == step.Register {
		return
	}

	if rightStep.Register == step.Register {
		instr.Left, instr.Right = instr.Right, instr.Left
		return
	}

	if step.Register == f.CPU.Call.Out[0] {
		_, rightIsCall := rightStep.Value.(*ssa.Call)

		if rightIsCall {
			_, leftIsCall := leftStep.Value.(*ssa.Call)

			if !leftIsCall || rightStep.Index > leftStep.Index {
				instr.Left, instr.Right = instr.Right, instr.Left
				return
			}
		}
	}

	_, leftIsInt := leftStep.Value.(*ssa.Int)
	_, rightIsInt := rightStep.Value.(*ssa.Int)

	if leftIsInt && !rightIsInt {
		instr.Left, instr.Right = instr.Right, instr.Left
		return
	}
}
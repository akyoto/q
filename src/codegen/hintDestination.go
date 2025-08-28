package codegen

import (
	"git.urbach.dev/cli/q/src/ssa"
)

// hintDestination recommends the destination register to its inputs.
func (f *Function) hintDestination(step *Step) {
	switch instr := step.Value.(type) {
	case *ssa.BinaryOp:
		if instr.Op.IsComparison() {
			return
		}

		// For operations that allow reordering, reorder the operands
		// so that the left operator matches the destination register.
		f.reorderOperands(step, instr)

		// For x86-64 it is advantageous to use the destination register
		// as the register for the left operand.
		f.ValueToStep[instr.Left].hint(step.Register)
	case *ssa.Phi:
		for _, variant := range instr.Arguments {
			variant := f.ValueToStep[variant]
			variant.Phis.Add(step)
			variant.hint(step.Register)
		}
	}
}
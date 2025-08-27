package codegen

import (
	"git.urbach.dev/cli/q/src/ssa"
)

// createHints recommends registers that a value must reside in later on.
func (f *Function) createHints(step *Step) {
	switch instr := step.Value.(type) {
	case *ssa.BinaryOp:
		if instr.Op.IsComparison() {
			return
		}

		if step.Register == -1 {
			return
		}

		// For operations that allow reordering, reorder the operands
		// so that the left operator matches the destination register.
		f.reorderOperands(step, instr)

		// For x86-64 it is advantageous to use the destination register
		// as the register for the left operand.
		f.ValueToStep[instr.Left].hint(step.Register)

	case *ssa.Call:
		for paramIndex, param := range instr.Arguments {
			f.ValueToStep[param].hint(f.CPU.Call.In[paramIndex])
		}

	case *ssa.CallExtern:
		for r, param := range instr.Arguments {
			if r >= len(f.CPU.ExternCall.In) {
				// Temporary hack to allow arguments 5 and 6 to be hinted as r10 and r11, then pushed later.
				if 1+r < len(f.CPU.ExternCall.Clobbered) {
					f.ValueToStep[param].hint(f.CPU.ExternCall.Clobbered[1+r])
				}

				continue
			}

			f.ValueToStep[param].hint(f.CPU.ExternCall.In[r])
		}

	case *ssa.FromTuple:
		if step.Register == -1 {
			step.Register = f.CPU.Call.Out[instr.Index]
		}

	case *ssa.Parameter:
		if step.Register == -1 {
			step.Register = f.CPU.Call.In[instr.Index]
		}

	case *ssa.Phi:
		for _, variant := range instr.Arguments {
			variant := f.ValueToStep[variant]
			variant.Phis = append(variant.Phis, step)

			if step.Register != -1 {
				variant.hint(step.Register)
			}
		}

	case *ssa.Return:
		for r, param := range instr.Arguments {
			f.ValueToStep[param].hint(f.CPU.Call.Out[r])
		}

	case *ssa.Syscall:
		for r, param := range instr.Arguments {
			f.ValueToStep[param].hint(f.CPU.Syscall.In[r])
		}
	}
}
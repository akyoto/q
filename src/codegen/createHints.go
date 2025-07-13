package codegen

import "git.urbach.dev/cli/q/src/ssa"

// createHints recommends registers that a value must reside in later on.
func (f *Function) createHints(step *step) {
	switch instr := step.Value.(type) {
	case *ssa.Call:
		for paramIndex, param := range instr.Arguments {
			f.ValueToStep[param].Hint(f.CPU.Call.In[paramIndex])
		}

	case *ssa.CallExtern:
		for r, param := range instr.Arguments {
			if r >= len(f.CPU.ExternCall.In) {
				// Temporary hack to allow arguments 5 and 6 to be hinted as r10 and r11, then pushed later.
				f.ValueToStep[param].Hint(f.CPU.ExternCall.Clobbered[1+r])
				continue
			}

			f.ValueToStep[param].Hint(f.CPU.ExternCall.In[r])
		}

	case *ssa.Parameter:
		f.ValueToStep[instr].Register = f.CPU.Call.In[instr.Index]

	case *ssa.Return:
		for r, param := range instr.Arguments {
			f.ValueToStep[param].Hint(f.CPU.Call.Out[r])
		}

	case *ssa.Syscall:
		for r, param := range instr.Arguments {
			f.ValueToStep[param].Hint(f.CPU.Syscall.In[r])
		}
	}
}
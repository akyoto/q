package codegen

import (
	"git.urbach.dev/cli/q/src/ssa"
)

// hintABI recommends ABI registers that a value must reside in later on.
// These register hints have the highest priority.
func (f *Function) hintABI(step *Step) {
	switch instr := step.Value.(type) {
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

		for _, user := range step.Value.Users() {
			switch user := user.(type) {
			case *ssa.BinaryOp:
				if !user.Op.IsComparison() && user.Left == step.Value {
					f.ValueToStep[user].hint(step.Register)
				}
			case *ssa.Phi:
				f.ValueToStep[user].hint(step.Register)
			case *ssa.UnaryOp:
				f.ValueToStep[user].hint(step.Register)
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
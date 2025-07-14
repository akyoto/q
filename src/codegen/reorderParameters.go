package codegen

import "git.urbach.dev/cli/q/src/ssa"

// reorderParameters reorders the sequence of parameter initialization and assigns new registers if needed.
func (f *Function) reorderParameters() {
	usedRegisters := 0
	futureRegisters := 0

	for i, step := range f.Steps {
		param, isParam := step.Value.(*ssa.Parameter)

		if !isParam {
			break
		}

		currentRegister := f.CPU.Call.In[param.Index]

		if futureRegisters&(1<<currentRegister) != 0 {
			bringToFront(f.Steps[:i+1], i)

			for h := range i + 1 {
				f.Steps[h].Index = h
			}

			if usedRegisters&(1<<step.Register) != 0 {
				users := step.Value.Users()
				alive := f.ValueToStep[users[len(users)-1]].Index
				step.Register = f.findFreeRegister(f.Steps[:alive])
			}
		}

		usedRegisters |= (1 << currentRegister)
		futureRegisters |= (1 << step.Register)
	}
}
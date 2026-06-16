package codegen

import (
	"git.urbach.dev/cli/q/src/set"
	"git.urbach.dev/cli/q/src/ssa"
)

// reorderParameters reorders the sequence of parameter initialization and assigns new registers if needed.
func (f *Function) reorderParameters() {
	usedRegisters := bitSet(0)
	futureRegisters := bitSet(0)

	for i, step := range f.Steps {
		param, isParam := step.Value.(*ssa.Parameter)

		if !isParam {
			break
		}

		currentRegister := f.CPU.Call.In[param.Index]

		if futureRegisters.Has(currentRegister) {
			set.BringToFront(f.Steps[:i+1], i)

			for h := range i + 1 {
				f.Steps[h].Index = Index(h)
			}

			if usedRegisters.Has(step.Register) {
				f.assignFreeRegister(step)
			}
		}

		usedRegisters.Set(currentRegister)
		futureRegisters.Set(step.Register)
	}
}
package codegen

import (
	"slices"

	"git.urbach.dev/cli/q/src/ssa"
)

// addPreservedRegisters creates a list of registers that need to be preserved.
func (f *Function) addPreservedRegisters() {
	mustPreserve := f.CPU.Call.Preserved

	for _, step := range f.Steps {
		register, isRegister := step.Value.(*ssa.Register)

		if isRegister && step.Register == register.Register {
			continue
		}

		if slices.Contains(mustPreserve, step.Register) {
			f.Preserved.Add(step.Register)
		}
	}
}
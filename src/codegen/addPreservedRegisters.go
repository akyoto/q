package codegen

import "slices"

// addPreservedRegisters creates a list of registers that need to be preserved.
func (f *Function) addPreservedRegisters() {
	mustPreserve := f.CPU.Call.Preserved

	for _, step := range f.Steps {
		if slices.Contains(mustPreserve, step.Register) {
			f.Preserved.Add(step.Register)
		}
	}
}
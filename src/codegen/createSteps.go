package codegen

import (
	"slices"

	"git.urbach.dev/cli/q/src/ssa"
)

// createSteps builds a series of instructions from the SSA values in the IR.
func (f *Function) createSteps(ir ssa.IR) {
	count := ir.CountValues()
	storage := make([]step, count)
	f.Steps = make([]*step, count)
	f.ValueToStep = make(map[ssa.Value]*step, count)

	for i, instr := range ir.Values {
		step := &storage[i]
		step.Index = i
		step.Value = instr
		step.Register = -1
		f.Steps[i] = step
		f.ValueToStep[instr] = step
	}

	for _, step := range slices.Backward(f.Steps) {
		f.createHints(step)
		f.createLiveRanges(step)
	}

	f.reorderParameters()
	f.fixRegisterConflicts()
	f.addPreservedRegisters()
}
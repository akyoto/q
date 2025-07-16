package codegen

import (
	"slices"

	"git.urbach.dev/cli/q/src/ssa"
)

// createSteps builds a series of instructions from the SSA values in the IR.
func (f *Function) createSteps(ir ssa.IR) {
	count := ir.CountValues() + len(ir.Blocks) - 1
	storage := make([]step, count)
	f.Steps = make([]*step, count)
	f.ValueToStep = make(map[ssa.Value]*step, count)
	i := 0

	for _, block := range ir.Blocks {
		if block != ir.Blocks[0] {
			step := &storage[i]
			step.Index = i
			step.Value = &Label{Name: block.Label}
			step.Register = -1
			f.Steps[i] = step
			i++
		}

		for _, instr := range block.Instructions {
			step := &storage[i]
			step.Index = i
			step.Value = instr
			step.Register = -1
			f.Steps[i] = step
			f.ValueToStep[instr] = step
			i++
		}
	}

	for _, step := range slices.Backward(f.Steps) {
		f.createHints(step)
		f.createLiveRanges(step)
	}

	f.reorderParameters()
	f.fixRegisterConflicts()
	f.addPreservedRegisters()
}
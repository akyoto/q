package codegen

import (
	"slices"

	"git.urbach.dev/cli/q/src/ssa"
)

// createSteps builds a series of instructions from the SSA values in the IR.
func (f *Function) createSteps(ir ssa.IR) {
	count := ir.CountValues() + len(ir.Blocks) - 1
	storage := make([]Step, count)
	f.Steps = make([]*Step, count)
	f.ValueToStep = make(map[ssa.Value]*Step, count)
	i := 0

	for _, block := range ir.Blocks {
		if block != ir.Blocks[0] {
			step := &storage[i]
			step.Index = i
			step.Value = &Label{Name: block.Label}
			step.Block = block
			step.Register = -1
			f.Steps[i] = step
			i++
		}

		for _, instr := range block.Instructions {
			step := &storage[i]
			step.Index = i
			step.Value = instr
			step.Block = block
			step.Register = -1
			f.Steps[i] = step
			f.ValueToStep[instr] = step
			i++
		}
	}

	for _, step := range slices.Backward(f.Steps) {
		f.hintABI(step)
		f.createLiveRanges(step)
	}

	for _, step := range slices.Backward(f.Steps) {
		if step.Register == -1 && f.needsRegister(step) {
			f.assignFreeRegister(step)
		}

		f.hintDestination(step)
	}

	f.reorderParameters()
	f.fixRegisterConflicts()
	f.addPreservedRegisters()
}
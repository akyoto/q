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
	f.BlockToRegion = make(map[*ssa.Block]region, len(ir.Blocks))
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

		f.BlockToRegion[block] = region{
			Start: uint32(i),
			End:   uint32(i + len(block.Instructions)),
		}

		for _, instr := range block.Instructions {
			step := &storage[i]
			step.Index = i
			step.Value = instr
			step.Block = block
			step.Register = -1
			step.Live = make([]*Step, 0, 4)
			f.Steps[i] = step
			f.ValueToStep[instr] = step
			i++
		}
	}

	f.reorderPhis()

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
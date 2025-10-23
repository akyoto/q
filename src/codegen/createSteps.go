package codegen

import "git.urbach.dev/cli/q/src/ssa"

// createSteps builds a series of instructions from the SSA values in the IR.
func createSteps(ir ssa.IR) IR {
	count := ir.CountValues() + len(ir.Blocks) - 1
	storage := make([]Step, count)
	steps := make([]*Step, count)
	valueToStep := make(map[ssa.Value]*Step, count)
	blockToRegion := make(map[*ssa.Block]region, len(ir.Blocks))
	i := 0

	for _, block := range ir.Blocks {
		if block != ir.Blocks[0] {
			step := &storage[i]
			step.Index = i
			step.Value = &Label{Name: block.Label}
			step.Block = block
			step.Register = -1
			steps[i] = step
			i++
		}

		blockToRegion[block] = region{
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
			steps[i] = step
			valueToStep[instr] = step
			i++
		}
	}

	return IR{steps, valueToStep, blockToRegion}
}
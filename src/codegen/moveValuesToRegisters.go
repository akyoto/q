package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
)

// moveValuesToRegisters moves the values to the destination registers.
func (f *Function) moveValuesToRegisters(values []ssa.Value, registers []cpu.Register, live []*Step) {
	moves := make([]*asm.Move, 0, len(values))
	sourceSteps := map[cpu.Register]*Step{}

	for i, arg := range values {
		sourceStep := f.ValueToStep[arg]
		source := sourceStep.Register
		destination := registers[i]

		if f.isSpilled(source) {
			sourceSteps[source] = sourceStep
		}

		moves = append(moves, &asm.Move{
			Destination: destination,
			Source:      source,
		})
	}

	free := f.freeTempRegisters(live)
	scheduled, ok := ScheduleMoves(moves, free)

	if !ok {
		panic("no free register for move scheduling")
	}

	for _, move := range scheduled {
		if f.isSpilled(move.Source) {
			f.loadSpill(sourceSteps[move.Source], move.Destination)
			continue
		}

		f.Assembler.Append(move)
	}
}
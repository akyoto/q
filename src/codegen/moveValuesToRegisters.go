package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
)

// moveValuesToRegisters moves the values to the destination registers.
func (f *Function) moveValuesToRegisters(values []ssa.Value, registers []cpu.Register) {
	start := len(f.Assembler.Instructions)

	for i, arg := range values {
		sourceStep := f.ValueToStep[arg]
		source := sourceStep.Register
		destination := registers[i]

		if f.isSpilled(source) {
			f.loadSpill(sourceStep, destination)
			continue
		}

		if source == destination {
			continue
		}

		f.Assembler.Append(&asm.Move{
			Destination: destination,
			Source:      source,
		})
	}

	reorderMoves(f.Assembler.Instructions[start:])
}
package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/cpu"
)

// moveCallResult moves the result from the call output register to the destination.
func (f *Function) moveCallResult(step *Step, source cpu.Register) {
	destination := step.Register

	if destination == -1 || destination == source {
		return
	}

	isSpilled := f.isSpilled(destination)

	if isSpilled {
		f.storeSpill(step, source)
		return
	}

	f.Assembler.Append(&asm.Move{
		Destination: destination,
		Source:      source,
	})
}
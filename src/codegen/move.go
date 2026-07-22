package codegen

import "git.urbach.dev/cli/q/src/asm"

// move moves source register to destination register and handles virtual registers.
func (f *Function) move(destinationStep *Step, sourceStep *Step, step *Step) {
	source := sourceStep.Register
	destination := destinationStep.Register

	if source == destination {
		return
	}

	sourceIsSpilled := f.isSpilled(source)
	destinationIsSpilled := f.isSpilled(destination)

	switch {
	case !sourceIsSpilled && !destinationIsSpilled:
		f.Assembler.Append(&asm.Move{
			Destination: destination,
			Source:      source,
		})

	case sourceIsSpilled && !destinationIsSpilled:
		f.loadSpill(sourceStep, destination)

	case !sourceIsSpilled && destinationIsSpilled:
		f.storeSpill(destinationStep, source)

	case sourceIsSpilled && destinationIsSpilled:
		source = f.resolveOperand(sourceStep, step.Live)
		f.storeSpill(destinationStep, source)
	}
}
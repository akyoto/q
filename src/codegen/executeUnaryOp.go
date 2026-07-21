package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

func (f *Function) executeUnaryOp(step *Step, instr *ssa.UnaryOp) {
	operand := f.ValueToStep[instr.Operand]
	source := f.resolveOperand(operand, step.Live)
	destination := step.Register
	isSpilled := f.isSpilled(destination)

	if isSpilled {
		destination = f.findTempRegister(step.Live)
	}

	switch instr.Op {
	case token.Negate:
		f.Assembler.Append(&asm.Negate{
			Destination: destination,
			Source:      source,
		})

	case token.Not:
		f.Assembler.Append(&asm.CompareNumber{
			Destination: source,
			Number:      0,
		})

		f.conditionalSet(destination, token.Equal, false)

	default:
		panic("not implemented: " + instr.String())
	}

	if isSpilled {
		f.storeSpill(step, destination)
	}
}
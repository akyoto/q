package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

func (f *Function) executeUnaryOp(step *Step, instr *ssa.UnaryOp) {
	operand := f.ValueToStep[instr.Operand]

	switch instr.Op {
	case token.Negate:
		f.Assembler.Append(&asm.Negate{
			Destination: step.Register,
			Source:      operand.Register,
		})

	case token.Not:
		f.Assembler.Append(&asm.CompareNumber{
			Destination: operand.Register,
			Number:      0,
		})

		f.conditionalSet(step.Register, token.Equal)

	default:
		panic("not implemented: " + instr.String())
	}
}
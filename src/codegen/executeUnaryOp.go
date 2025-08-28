package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

func (f *Function) executeUnaryOp(step *Step, instr *ssa.UnaryOp) {
	left := f.ValueToStep[instr.Operand]

	switch instr.Op {
	case token.Negate:
		f.Assembler.Append(&asm.Negate{
			Destination: step.Register,
			Source:      left.Register,
		})

	case token.Not:
		panic("not implemented: logical not")

	default:
		panic("not implemented: " + instr.String())
	}
}
package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/config"
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
		f.Assembler.Append(&asm.CompareNumber{
			Destination: left.Register,
			Number:      0,
		})

		if f.build.Arch == config.X86 {
			f.Assembler.Append(&asm.MoveNumber{
				Destination: step.Register,
				Number:      0,
			})
		}

		f.Assembler.Append(&asm.ConditionalSet{
			Destination: step.Register,
			Condition:   token.Equal,
		})

	default:
		panic("not implemented: " + instr.String())
	}
}
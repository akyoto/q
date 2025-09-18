package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

func (f *Function) executeBinaryOp(step *Step, instr *ssa.BinaryOp) {
	left := f.ValueToStep[instr.Left]
	right := f.ValueToStep[instr.Right]

	if instr.Op.IsComparison() {
		number, isInt := right.Value.(*ssa.Int)

		if isInt && right.Register == -1 {
			f.Assembler.Append(&asm.CompareNumber{Destination: left.Register, Number: number.Int})
		} else {
			f.Assembler.Append(&asm.Compare{Destination: left.Register, Source: right.Register})
		}

		return
	}

	number, isInt := right.Value.(*ssa.Int)

	if isInt && right.Register == -1 {
		switch instr.Op {
		case token.Add:
			f.Assembler.Append(&asm.AddNumber{
				Destination: step.Register,
				Source:      left.Register,
				Number:      number.Int,
			})

		case token.And:
			f.Assembler.Append(&asm.AndNumber{
				Destination: step.Register,
				Source:      left.Register,
				Number:      number.Int,
			})

		case token.Or:
			f.Assembler.Append(&asm.OrNumber{
				Destination: step.Register,
				Source:      left.Register,
				Number:      number.Int,
			})

		case token.Shl:
			f.Assembler.Append(&asm.ShiftLeftNumber{
				Destination: step.Register,
				Source:      left.Register,
				Number:      number.Int,
			})

		case token.Shr:
			if types.IsUnsigned(left.Value.Type()) {
				f.Assembler.Append(&asm.ShiftRightNumber{
					Destination: step.Register,
					Source:      left.Register,
					Number:      number.Int,
				})
			} else {
				f.Assembler.Append(&asm.ShiftRightSignedNumber{
					Destination: step.Register,
					Source:      left.Register,
					Number:      number.Int,
				})
			}

		case token.Sub:
			f.Assembler.Append(&asm.SubtractNumber{
				Destination: step.Register,
				Source:      left.Register,
				Number:      number.Int,
			})

		case token.Xor:
			f.Assembler.Append(&asm.XorNumber{
				Destination: step.Register,
				Source:      left.Register,
				Number:      number.Int,
			})

		default:
			panic("not implemented: " + instr.String())
		}

		return
	}

	switch instr.Op {
	case token.Add:
		f.Assembler.Append(&asm.Add{
			Destination: step.Register,
			Source:      left.Register,
			Operand:     right.Register,
		})

	case token.Div:
		f.Assembler.Append(&asm.Divide{
			Destination: step.Register,
			Source:      left.Register,
			Operand:     right.Register,
		})

	case token.Mul:
		f.Assembler.Append(&asm.Multiply{
			Destination: step.Register,
			Source:      left.Register,
			Operand:     right.Register,
		})

	case token.Sub:
		f.Assembler.Append(&asm.Subtract{
			Destination: step.Register,
			Source:      left.Register,
			Operand:     right.Register,
		})

	case token.Mod:
		f.Assembler.Append(&asm.Modulo{
			Destination: step.Register,
			Source:      left.Register,
			Operand:     right.Register,
		})

	case token.And:
		f.Assembler.Append(&asm.And{
			Destination: step.Register,
			Source:      left.Register,
			Operand:     right.Register,
		})

	case token.Or:
		f.Assembler.Append(&asm.Or{
			Destination: step.Register,
			Source:      left.Register,
			Operand:     right.Register,
		})

	case token.Xor:
		f.Assembler.Append(&asm.Xor{
			Destination: step.Register,
			Source:      left.Register,
			Operand:     right.Register,
		})

	case token.Shl:
		f.Assembler.Append(&asm.ShiftLeft{
			Destination: step.Register,
			Source:      left.Register,
			Operand:     right.Register,
		})

	case token.Shr:
		f.Assembler.Append(&asm.ShiftRightSigned{
			Destination: step.Register,
			Source:      left.Register,
			Operand:     right.Register,
		})

	default:
		panic("not implemented: " + instr.String())
	}
}
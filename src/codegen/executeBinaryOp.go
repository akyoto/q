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
	source := f.resolveOperand(left, step.Live)
	operand := f.resolveOperand(right, step.Live)
	destination := step.Register
	isSpilled := f.isSpilled(destination)

	if isSpilled {
		destination = f.findTempRegister(step.Live)
	}

	if instr.Op.IsComparison() {
		number, isInt := right.Value.(*ssa.Int)

		if isInt && right.Register == -1 {
			f.Assembler.Append(&asm.CompareNumber{Destination: source, Number: number.Int})
		} else {
			f.Assembler.Append(&asm.Compare{Destination: source, Source: operand})
		}

		if step.Register != -1 {
			unsigned := types.IsUnsigned(left.Value.Type()) || types.IsUnsigned(right.Value.Type())
			f.conditionalSet(destination, instr.Op, unsigned)

			if isSpilled {
				f.storeSpill(step, destination)
			}
		}

		return
	}

	number, isInt := right.Value.(*ssa.Int)

	if isInt && right.Register == -1 {
		switch instr.Op {
		case token.Add:
			f.Assembler.Append(&asm.AddNumber{
				Destination: destination,
				Source:      source,
				Number:      number.Int,
			})

		case token.And:
			f.Assembler.Append(&asm.AndNumber{
				Destination: destination,
				Source:      source,
				Number:      number.Int,
			})

		case token.Or:
			f.Assembler.Append(&asm.OrNumber{
				Destination: destination,
				Source:      source,
				Number:      number.Int,
			})

		case token.Shl:
			f.Assembler.Append(&asm.ShiftLeftNumber{
				Destination: destination,
				Source:      source,
				Number:      number.Int,
			})

		case token.Shr:
			if types.IsUnsigned(left.Value.Type()) {
				f.Assembler.Append(&asm.ShiftRightNumber{
					Destination: destination,
					Source:      source,
					Number:      number.Int,
				})
			} else {
				f.Assembler.Append(&asm.ShiftRightSignedNumber{
					Destination: destination,
					Source:      source,
					Number:      number.Int,
				})
			}

		case token.Sub:
			f.Assembler.Append(&asm.SubtractNumber{
				Destination: destination,
				Source:      source,
				Number:      number.Int,
			})

		case token.Xor:
			f.Assembler.Append(&asm.XorNumber{
				Destination: destination,
				Source:      source,
				Number:      number.Int,
			})

		default:
			panic("not implemented: " + instr.String())
		}

		if isSpilled {
			f.storeSpill(step, destination)
		}

		return
	}

	switch instr.Op {
	case token.Add:
		f.Assembler.Append(&asm.Add{
			Destination: destination,
			Source:      source,
			Operand:     operand,
		})

	case token.Div:
		if types.IsUnsigned(left.Value.Type()) {
			f.Assembler.Append(&asm.Divide{
				Destination: destination,
				Source:      source,
				Operand:     operand,
			})
		} else {
			f.Assembler.Append(&asm.DivideSigned{
				Destination: destination,
				Source:      source,
				Operand:     operand,
			})
		}

	case token.Mul:
		f.Assembler.Append(&asm.Multiply{
			Destination: destination,
			Source:      source,
			Operand:     operand,
		})

	case token.Sub:
		f.Assembler.Append(&asm.Subtract{
			Destination: destination,
			Source:      source,
			Operand:     operand,
		})

	case token.Mod:
		if types.IsUnsigned(left.Value.Type()) {
			f.Assembler.Append(&asm.Modulo{
				Destination: destination,
				Source:      source,
				Operand:     operand,
			})
		} else {
			f.Assembler.Append(&asm.ModuloSigned{
				Destination: destination,
				Source:      source,
				Operand:     operand,
			})
		}

	case token.And:
		f.Assembler.Append(&asm.And{
			Destination: destination,
			Source:      source,
			Operand:     operand,
		})

	case token.Or:
		f.Assembler.Append(&asm.Or{
			Destination: destination,
			Source:      source,
			Operand:     operand,
		})

	case token.Xor:
		f.Assembler.Append(&asm.Xor{
			Destination: destination,
			Source:      source,
			Operand:     operand,
		})

	case token.Shl:
		f.Assembler.Append(&asm.ShiftLeft{
			Destination: destination,
			Source:      source,
			Operand:     operand,
		})

	case token.Shr:
		if types.IsUnsigned(left.Value.Type()) {
			f.Assembler.Append(&asm.ShiftRight{
				Destination: destination,
				Source:      source,
				Operand:     operand,
			})
		} else {
			f.Assembler.Append(&asm.ShiftRightSigned{
				Destination: destination,
				Source:      source,
				Operand:     operand,
			})
		}

	default:
		panic("not implemented: " + instr.String())
	}

	if isSpilled {
		f.storeSpill(step, destination)
	}
}
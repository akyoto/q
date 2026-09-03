package codegen

import (
	"slices"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

func (f *Function) executeBinaryOp(step *Step, instr *ssa.BinaryOp) {
	if step.Register == -1 && !instr.Op.IsComparison() {
		return
	}

	left := f.ValueToStep[instr.Left]
	right := f.ValueToStep[instr.Right]
	live := slices.Concat(step.Live, []*Step{left, right})
	source := f.resolveOperand(left, live)
	avoid := []cpu.Register{source}

	if instr.Op == token.Div || instr.Op == token.Mod {
		avoid = append(avoid, f.CPU.DivisorRestricted...)
	}

	operand := f.resolveOperand(right, live, avoid...)
	destination := step.Register
	isSpilled := f.isSpilled(destination)

	if isSpilled {
		destination = f.findTempRegister(live, source, operand)
	}

	if instr.Op.IsComparison() {
		f.emitComparison(step, left, right, source, operand, destination, instr.Op)
	} else if isImmediate(right) {
		f.emitArithmeticImmediate(left, source, destination, right.Value.(*ssa.Int).Int, instr.Op)
	} else {
		f.emitArithmeticRegister(left, source, operand, destination, instr.Op)
	}

	if isSpilled {
		f.storeSpill(step, destination)
	}
}

// emitComparison emits a compare followed by a conditional set when the result register is needed.
func (f *Function) emitComparison(step *Step, left *Step, right *Step, source cpu.Register, operand cpu.Register, destination cpu.Register, op token.Kind) {
	if isImmediate(right) {
		f.Assembler.Append(&asm.CompareNumber{Destination: source, Number: right.Value.(*ssa.Int).Int})
	} else {
		f.Assembler.Append(&asm.Compare{Destination: source, Source: operand})
	}

	if step.Register != -1 {
		unsigned := types.IsUnsigned(left.Value.Type()) || types.IsUnsigned(right.Value.Type())
		f.conditionalSet(destination, op, unsigned)
	}
}

// emitArithmeticImmediate emits an arithmetic instruction with an immediate right operand.
func (f *Function) emitArithmeticImmediate(left *Step, source cpu.Register, destination cpu.Register, number int, op token.Kind) {
	switch op {
	case token.Add:
		f.Assembler.Append(&asm.AddNumber{Destination: destination, Source: source, Number: number})
	case token.And, token.LogicalAnd:
		f.Assembler.Append(&asm.AndNumber{Destination: destination, Source: source, Number: number})
	case token.Or, token.LogicalOr:
		f.Assembler.Append(&asm.OrNumber{Destination: destination, Source: source, Number: number})
	case token.Sub:
		f.Assembler.Append(&asm.SubtractNumber{Destination: destination, Source: source, Number: number})
	case token.Xor:
		f.Assembler.Append(&asm.XorNumber{Destination: destination, Source: source, Number: number})
	case token.Shl:
		f.Assembler.Append(&asm.ShiftLeftNumber{Destination: destination, Source: source, Number: number})
	case token.Shr:
		f.emitShiftRightImmediate(source, destination, number, left.Value.Type())
	default:
		panic("not implemented: " + op.String())
	}
}

// emitArithmeticRegister emits an arithmetic instruction with a register right operand.
func (f *Function) emitArithmeticRegister(left *Step, source cpu.Register, operand cpu.Register, destination cpu.Register, op token.Kind) {
	switch op {
	case token.Add:
		f.Assembler.Append(&asm.Add{Destination: destination, Source: source, Operand: operand})
	case token.Sub:
		f.Assembler.Append(&asm.Subtract{Destination: destination, Source: source, Operand: operand})
	case token.Mul:
		f.Assembler.Append(&asm.Multiply{Destination: destination, Source: source, Operand: operand})
	case token.And, token.LogicalAnd:
		f.Assembler.Append(&asm.And{Destination: destination, Source: source, Operand: operand})
	case token.Or, token.LogicalOr:
		f.Assembler.Append(&asm.Or{Destination: destination, Source: source, Operand: operand})
	case token.Xor:
		f.Assembler.Append(&asm.Xor{Destination: destination, Source: source, Operand: operand})
	case token.Shl:
		f.Assembler.Append(&asm.ShiftLeft{Destination: destination, Source: source, Operand: operand})
	case token.Shr:
		f.emitShiftRightRegister(source, operand, destination, left.Value.Type())
	case token.Div:
		f.emitDivision(source, operand, destination, left.Value.Type())
	case token.Mod:
		f.emitModulo(source, operand, destination, left.Value.Type())
	default:
		panic("not implemented: " + op.String())
	}
}

// emitShiftRightImmediate emits a shift-right instruction with an immediate count.
func (f *Function) emitShiftRightImmediate(source cpu.Register, destination cpu.Register, count int, typ types.Type) {
	if types.IsUnsigned(typ) {
		f.Assembler.Append(&asm.ShiftRightNumber{Destination: destination, Source: source, Number: count})
	} else {
		f.Assembler.Append(&asm.ShiftRightSignedNumber{Destination: destination, Source: source, Number: count})
	}
}

// emitShiftRightRegister emits a shift-right instruction with a register count.
func (f *Function) emitShiftRightRegister(source cpu.Register, operand cpu.Register, destination cpu.Register, typ types.Type) {
	if types.IsUnsigned(typ) {
		f.Assembler.Append(&asm.ShiftRight{Destination: destination, Source: source, Operand: operand})
	} else {
		f.Assembler.Append(&asm.ShiftRightSigned{Destination: destination, Source: source, Operand: operand})
	}
}

// emitDivision emits a division instruction.
func (f *Function) emitDivision(source cpu.Register, operand cpu.Register, destination cpu.Register, typ types.Type) {
	if types.IsUnsigned(typ) {
		f.Assembler.Append(&asm.Divide{Destination: destination, Source: source, Operand: operand})
	} else {
		f.Assembler.Append(&asm.DivideSigned{Destination: destination, Source: source, Operand: operand})
	}
}

// emitModulo emits a modulo instruction.
func (f *Function) emitModulo(source cpu.Register, operand cpu.Register, destination cpu.Register, typ types.Type) {
	if types.IsUnsigned(typ) {
		f.Assembler.Append(&asm.Modulo{Destination: destination, Source: source, Operand: operand})
	} else {
		f.Assembler.Append(&asm.ModuloSigned{Destination: destination, Source: source, Operand: operand})
	}
}

// isImmediate returns true if the operand is an integer with no assigned register.
func isImmediate(operand *Step) bool {
	_, isInt := operand.Value.(*ssa.Int)
	return isInt && operand.Register == -1
}
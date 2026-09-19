package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
	"git.urbach.dev/cli/q/src/x86"
)

type archX86 struct {
	build *config.Build
}

func (a archX86) canEncodeNumber(instr ssa.Value, number *ssa.Int) bool {
	switch instr := instr.(type) {
	case *ssa.BinaryOp:
		if number != instr.Right {
			return false
		}

		if instr.Op.IsComparison() {
			return cpu.SizeInt(number.Int) <= 4
		}

		switch instr.Op {
		case token.Add, token.And, token.Or, token.Sub, token.Xor:
			return cpu.SizeInt(number.Int) <= 4
		case token.Shl, token.Shr:
			return number.Int >= 0 && number.Int <= 63
		}
	case *ssa.Load:
		if instr.Memory.Index != number {
			return false
		}

		if instr.Memory.Scale {
			return false
		}

		return number.Int >= -128 && number.Int <= 127
	case *ssa.Store:
		if instr.Memory.Scale {
			return false
		}

		if instr.Value == number && cpu.SizeInt(number.Int) <= 4 {
			return true
		}

		return instr.Memory.Index == number && cpu.SizeInt(number.Int) <= 1
	}

	return false
}

func (a archX86) canStoreNumber(typ types.Type, number int) bool {
	unsigned := types.IsUnsigned(typ)
	return !unsigned && cpu.SizeInt(number) <= 4 || unsigned && cpu.SizeUint(number) <= 4
}

func (a archX86) casOldValue(f *Function, register cpu.Register) cpu.Register {
	if register != x86.R0 {
		f.Assembler.Append(&asm.Move{
			Destination: x86.R0,
			Source:      register,
		})
	}

	return x86.R0
}

func (a archX86) casResultRegister(register cpu.Register) cpu.Register {
	return x86.R0
}

func (a archX86) conditionalSet(f *Function, register cpu.Register, op token.Kind, unsigned bool) {
	f.Assembler.Append(&asm.MoveNumber{
		Destination: register,
		Number:      0,
	})

	f.Assembler.Append(&asm.ConditionalSet{
		Destination: register,
		Condition:   tokenToCondition(op, unsigned),
	})
}

func (a archX86) conflictsWithStackPointer(register cpu.Register) bool {
	return false
}

func (a archX86) loadTLS(f *Function, destination cpu.Register, label string) {
	switch a.build.OS {
	case config.Linux:
		f.Assembler.Append(&asm.ReadSystemRegister{
			Destination:    destination,
			SystemRegister: x86.FS,
		})
	case config.Windows:
		f.Assembler.Append(&asm.ReadSystemRegister{
			Destination:    destination,
			SystemRegister: x86.GS,
		})

		f.Assembler.Append(&asm.AddNumber{
			Destination: destination,
			Source:      destination,
			Number:      0x1000 + WindowsTLSOffset + WindowsTLSSize - 0x20,
		})
	default:
		f.Assembler.Append(&asm.MoveLabel{
			Destination: destination,
			Label:       label,
		})
	}
}

func (a archX86) operandConflict(instr ssa.Value, value ssa.Value) bool {
	binaryOp, isBinaryOp := instr.(*ssa.BinaryOp)
	return isBinaryOp && !binaryOp.Op.IsComparison() && value == binaryOp.Right
}
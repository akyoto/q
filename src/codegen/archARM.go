package codegen

import (
	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

type archARM struct {
	build *config.Build
}

func (a archARM) canEncodeNumber(instr ssa.Value, number *ssa.Int) bool {
	switch instr := instr.(type) {
	case *ssa.BinaryOp:
		if number != instr.Right {
			return false
		}

		if instr.Op.IsComparison() {
			_, encodable := arm.CompareRegisterNumber(0, number.Int)
			return encodable
		}

		switch instr.Op {
		case token.Add:
			_, encodable := arm.AddRegisterNumber(0, 0, number.Int)
			return encodable
		case token.And:
			_, encodable := arm.AndRegisterNumber(0, 0, number.Int)
			return encodable
		case token.Or:
			_, encodable := arm.OrRegisterNumber(0, 0, number.Int)
			return encodable
		case token.Shl, token.Shr:
			return number.Int >= 0 && number.Int <= 63
		case token.Sub:
			_, encodable := arm.SubRegisterNumber(0, 0, number.Int)
			return encodable
		case token.Xor:
			_, encodable := arm.XorRegisterNumber(0, 0, number.Int)
			return encodable
		}
	case *ssa.Load:
		if instr.Memory.Index != number {
			return false
		}

		if instr.Memory.Scale {
			return number.Int >= 0 && number.Int <= 4095
		}

		return number.Int >= -256 && number.Int <= 255
	case *ssa.Store:
		if instr.Memory.Index != number {
			return false
		}

		if instr.Memory.Scale {
			return number.Int >= 0 && number.Int <= 4095
		}

		return number.Int >= -256 && number.Int <= 255
	}

	return false
}

func (a archARM) canStoreNumber(typ types.Type, number int) bool {
	return false
}

func (a archARM) casOldValue(f *Function, register cpu.Register) cpu.Register {
	return register
}

func (a archARM) casResultRegister(register cpu.Register) cpu.Register {
	return register
}

func (a archARM) conditionalSet(f *Function, register cpu.Register, op token.Kind, unsigned bool) {
	f.Assembler.Append(&asm.ConditionalSet{
		Destination: register,
		Condition:   tokenToCondition(op, unsigned),
	})
}

func (a archARM) conflictsWithStackPointer(register cpu.Register) bool {
	return register == arm.SP
}

func (a archARM) loadTLS(f *Function, destination cpu.Register, label string) {
	switch a.build.OS {
	case config.Linux:
		f.Assembler.Append(&asm.ReadSystemRegister{
			Destination:    destination,
			SystemRegister: arm.TPIDR_EL0,
		})
	case config.Windows:
		f.Assembler.Append(&asm.AddNumber{
			Destination: destination,
			Source:      arm.X18,
			Number:      0x1000,
		})

		f.Assembler.Append(&asm.AddNumber{
			Destination: destination,
			Source:      destination,
			Number:      WindowsTLSOffset + WindowsTLSSize - 0x20,
		})
	default:
		f.Assembler.Append(&asm.MoveLabel{
			Destination: destination,
			Label:       label,
		})
	}
}

func (a archARM) operandConflict(instr ssa.Value, value ssa.Value) bool {
	binaryOp, isBinaryOp := instr.(*ssa.BinaryOp)
	return isBinaryOp && binaryOp.Op == token.Mod && (value == binaryOp.Left || value == binaryOp.Right)
}
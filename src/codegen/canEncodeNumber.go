package codegen

import (
	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// canEncodeNumber returns true if the architecture can encode an immediate for the given instruction.
func (f *Function) canEncodeNumber(instr ssa.Value, number *ssa.Int) bool {
	switch instr := instr.(type) {
	case *ssa.BinaryOp:
		if number != instr.Right {
			return false
		}

		switch f.build.Arch {
		case config.ARM:
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

		case config.X86:
			if instr.Op.IsComparison() {
				return cpu.SizeInt(number.Int) <= 4
			}

			switch instr.Op {
			case token.Add, token.And, token.Or, token.Sub, token.Xor:
				return cpu.SizeInt(number.Int) <= 4

			case token.Shl, token.Shr:
				return number.Int >= 0 && number.Int <= 63
			}
		}

	case *ssa.Load:
		if instr.Memory.Index != number {
			return false
		}

		switch f.build.Arch {
		case config.ARM:
			if instr.Memory.Scale {
				return number.Int >= 0 && number.Int <= 4095
			} else {
				return number.Int >= -256 && number.Int <= 255
			}
		case config.X86:
			if instr.Memory.Scale {
				return false
			}

			return number.Int >= -128 && number.Int <= 127
		}

	case *ssa.Store:
		switch f.build.Arch {
		case config.ARM:
			return false
		case config.X86:
			return instr.Value == number && cpu.SizeInt(number.Int) <= 4
		}
	}

	return false
}
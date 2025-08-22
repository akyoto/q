package codegen

import (
	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/sizeof"
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

			case token.Shl, token.Shr:
				return number.Int >= 0 && number.Int <= 63

			case token.Sub:
				_, encodable := arm.SubRegisterNumber(0, 0, number.Int)
				return encodable
			}

		case config.X86:
			if instr.Op.IsComparison() {
				return sizeof.Signed(number.Int) <= 4
			}

			switch instr.Op {
			case token.Add, token.Sub:
				return sizeof.Signed(number.Int) <= 4

			case token.Shl, token.Shr:
				return number.Int >= 0 && number.Int <= 63
			}
		}

	case *ssa.Store:
		switch f.build.Arch {
		case config.ARM:
			return false
		case config.X86:
			return instr.Value == number && sizeof.Signed(number.Int) <= 4
		}
	}

	return false
}
package codegen

import (
	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/sizeof"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// canEncodeNumber returns true if the architecture can encode an immediate for the given instruction.
func (f *Function) canEncodeNumber(instr *Step, value *ssa.Int) bool {
	binaryOp, isBinaryOp := instr.Value.(*ssa.BinaryOp)

	if !isBinaryOp {
		return false
	}

	if value != binaryOp.Right {
		return false
	}

	switch f.build.Arch {
	case config.ARM:
		if binaryOp.Op.IsComparison() {
			_, encodable := arm.CompareRegisterNumber(0, value.Int)
			return encodable
		}

		switch binaryOp.Op {
		case token.Add:
			_, encodable := arm.AddRegisterNumber(0, 0, value.Int)
			return encodable

		case token.Sub:
			_, encodable := arm.SubRegisterNumber(0, 0, value.Int)
			return encodable
		}

	case config.X86:
		if binaryOp.Op.IsComparison() {
			return sizeof.Signed(value.Int) <= 4
		}

		switch binaryOp.Op {
		case token.Add, token.Sub:
			return sizeof.Signed(value.Int) <= 4
		}
	}

	return false
}
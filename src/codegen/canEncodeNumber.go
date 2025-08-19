package codegen

import (
	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/sizeof"
	"git.urbach.dev/cli/q/src/ssa"
)

// canEncodeNumber returns true if the architecture can encode an immediate for the given instruction.
func (f *Function) canEncodeNumber(instr *Step, value *ssa.Int) bool {
	binaryOp, isBinaryOp := instr.Value.(*ssa.BinaryOp)

	if isBinaryOp {
		if binaryOp.Op.IsComparison() && value == binaryOp.Right {
			switch f.build.Arch {
			case config.ARM:
				_, encodable := arm.CompareRegisterNumber(instr.Register, value.Int)
				return encodable
			case config.X86:
				return sizeof.Signed(value.Int) <= 4
			}
		}
	}

	return false
}
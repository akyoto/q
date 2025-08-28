package codegen

import (
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executePhi(instr *ssa.Phi) {
	// Phi does not generate any machine instructions.
}
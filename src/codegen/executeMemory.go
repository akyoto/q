package codegen

import (
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeMemory(step *Step, instr *ssa.Memory) {
	// Memory does not generate any machine instructions.
}
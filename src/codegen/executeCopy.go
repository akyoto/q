package codegen

import (
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeCopy(step *Step, instr *ssa.Copy) {
	if step.Register == -1 {
		return
	}

	copy := f.ValueToStep[instr.Value]
	f.move(step, copy, step)
}
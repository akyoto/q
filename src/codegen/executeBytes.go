package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeBytes(step *Step, instr *ssa.Bytes) {
	f.Count.Data++
	label := f.CreateLabel("data", f.Count.Data)
	f.Assembler.Data.SetImmutable(label, instr.Bytes)

	f.Assembler.Append(&asm.MoveLabel{
		Destination: step.Register,
		Label:       label,
	})
}
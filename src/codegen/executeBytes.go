package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeBytes(step *Step, instr *ssa.Bytes) {
	destination := step.Register

	if destination == -1 {
		return
	}

	f.Count.Data++
	label := f.CreateLabel("data", f.Count.Data)
	f.Assembler.Data.SetImmutable(label, instr.Bytes)
	isSpilled := f.isSpilled(destination)

	if isSpilled {
		destination = f.findTempRegister(step.Live)
	}

	f.Assembler.Append(&asm.MoveLabel{
		Destination: destination,
		Label:       label,
	})

	if isSpilled {
		f.storeSpill(step, destination)
	}
}
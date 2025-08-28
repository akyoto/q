package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
)

func (f *Function) executeLabel(instr *Label) {
	f.Assembler.Append(&asm.Label{
		Name: instr.Name,
	})
}
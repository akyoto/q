package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
)

// jumpIfTrue jumps to the label if the previous comparison was true.
func (f *Function) jumpIfTrue(condition asm.Condition, label string) {
	f.Assembler.Append(&asm.Jump{Label: label, Condition: condition})
}
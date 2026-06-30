package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/token"
)

// jumpIfTrue jumps to the label if the previous comparison was true.
func (f *Function) jumpIfTrue(operator token.Kind, label string) {
	f.Assembler.Append(&asm.Jump{Label: label, Condition: operator})
}
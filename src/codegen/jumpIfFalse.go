package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/token"
)

// jumpIfFalse jumps to the label if the previous comparison was false.
func (f *Function) jumpIfFalse(operator token.Kind, label string) {
	switch operator {
	case token.Equal:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: token.NotEqual})
	case token.NotEqual:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: token.Equal})
	case token.Greater:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: token.LessEqual})
	case token.Less:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: token.GreaterEqual})
	case token.GreaterEqual:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: token.Less})
	case token.LessEqual:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: token.Greater})
	}
}
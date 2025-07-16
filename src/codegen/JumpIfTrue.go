package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/token"
)

// JumpIfTrue jumps to the label if the previous comparison was true.
func (f *Function) JumpIfTrue(operator token.Kind, label string) {
	switch operator {
	case token.Equal:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: token.Equal})
	case token.NotEqual:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: token.NotEqual})
	case token.Greater:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: token.Greater})
	case token.Less:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: token.Less})
	case token.GreaterEqual:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: token.GreaterEqual})
	case token.LessEqual:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: token.LessEqual})
	}
}
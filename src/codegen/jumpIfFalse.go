package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
)

// jumpIfFalse jumps to the label if the previous comparison was false.
func (f *Function) jumpIfFalse(condition asm.Condition, label string) {
	switch condition {
	case asm.Equal:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: asm.NotEqual})
	case asm.NotEqual:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: asm.Equal})
	case asm.Greater:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: asm.LessEqual})
	case asm.GreaterEqual:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: asm.Less})
	case asm.Less:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: asm.GreaterEqual})
	case asm.LessEqual:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: asm.Greater})
	case asm.UnsignedGreater:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: asm.UnsignedLessEqual})
	case asm.UnsignedGreaterEqual:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: asm.UnsignedLess})
	case asm.UnsignedLess:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: asm.UnsignedGreaterEqual})
	case asm.UnsignedLessEqual:
		f.Assembler.Append(&asm.Jump{Label: label, Condition: asm.UnsignedGreater})
	}
}
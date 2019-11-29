package assembler

import (
	"fmt"

	"github.com/akyoto/asm"
)

// label is used for instructions requiring a label.
type label struct {
	Mnemonic string
	Label    string
	size     byte
}

// Exec writes the instruction to the final assembler.
func (instr *label) Exec(a *asm.Assembler) {
	start := a.Len()

	switch instr.Mnemonic {
	case CALL:
		a.Call(instr.Label)

	case JE:
		a.JumpIfEqual(instr.Label)

	case JNE:
		a.JumpIfNotEqual(instr.Label)

	case JG:
		a.JumpIfGreater(instr.Label)

	case JGE:
		a.JumpIfGreaterOrEqual(instr.Label)

	case JL:
		a.JumpIfLess(instr.Label)

	case JLE:
		a.JumpIfLessOrEqual(instr.Label)

	case JMP:
		a.Jump(instr.Label)
	}

	instr.size = byte(a.Len() - start)
}

// Name returns the mnemonic.
func (instr *label) Name() string {
	return instr.Mnemonic
}

// Size returns the number of bytes consumed for the instruction.
func (instr *label) Size() byte {
	return instr.size
}

// String implements the string serialization.
func (instr *label) String() string {
	return fmt.Sprintf("[%d] %s %s", instr.size, instr.Mnemonic, instr.Label)
}

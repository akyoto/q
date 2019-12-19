package instructions

import (
	"fmt"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/assembler/mnemonics"
)

// Jump is used for instructions requiring a label.
type Jump struct {
	Base
	Label string
}

// Exec writes the instruction to the final assembler.
func (instr *Jump) Exec(a *asm.Assembler) {
	start := a.Len()

	switch instr.Mnemonic {
	case mnemonics.CALL:
		a.Call(instr.Label)

	case mnemonics.JE:
		a.JumpIfEqual(instr.Label)

	case mnemonics.JNE:
		a.JumpIfNotEqual(instr.Label)

	case mnemonics.JG:
		a.JumpIfGreater(instr.Label)

	case mnemonics.JGE:
		a.JumpIfGreaterOrEqual(instr.Label)

	case mnemonics.JL:
		a.JumpIfLess(instr.Label)

	case mnemonics.JLE:
		a.JumpIfLessOrEqual(instr.Label)

	case mnemonics.JMP:
		a.Jump(instr.Label)
	}

	instr.size = byte(a.Len() - start)
}

// String implements the string serialization.
func (instr *Jump) String() string {
	return fmt.Sprintf("[%d]   %s %s", instr.size, mnemonicColor.Sprint(instr.Mnemonic), instr.Label)
}

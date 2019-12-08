package assembler

import (
	"fmt"

	"github.com/akyoto/asm"
)

// standalone is used for instructions with no operands.
type standalone struct {
	Mnemonic string
	size     byte
}

// Exec writes the instruction to the final assembler.
func (instr *standalone) Exec(a *asm.Assembler) {
	start := a.Len()

	switch instr.Mnemonic {
	case RET:
		a.Return()

	case SYSCALL:
		a.Syscall()

	case CPUID:
		a.CPUID()
	}

	instr.size = byte(a.Len() - start)
}

// Name returns the mnemonic.
func (instr *standalone) Name() string {
	return instr.Mnemonic
}

// Size returns the number of bytes consumed for the instruction.
func (instr *standalone) Size() byte {
	return instr.size
}

// String implements the string serialization.
func (instr *standalone) String() string {
	return fmt.Sprintf("[%d]   %s", instr.size, mnemonicColor.Sprint(instr.Mnemonic))
}

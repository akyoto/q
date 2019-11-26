package assembler

import (
	"github.com/akyoto/asm"
)

// standalone is used for instructions with no operands.
type standalone struct {
	Mnemonic string
}

// Exec writes the instruction to the final assembler.
func (instr *standalone) Exec(a *asm.Assembler) {
	switch instr.Mnemonic {
	case RET:
		a.Return()

	case SYSCALL:
		a.Syscall()

	case CPUID:
		a.CPUID()
	}
}

// Name returns the mnemonic.
func (instr *standalone) Name() string {
	return instr.Mnemonic
}

// String implements the string serialization.
func (instr *standalone) String() string {
	return instr.Mnemonic
}

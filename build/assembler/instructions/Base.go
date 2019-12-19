package instructions

import (
	"fmt"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/assembler/mnemonics"
)

// Base represents the data that is common among all instructions.
type Base struct {
	Mnemonic string
	size     byte
}

// Name returns the mnemonic.
func (instr *Base) Name() string {
	return instr.Mnemonic
}

// SetName sets the mnemonic.
func (instr *Base) SetName(mnemonic string) {
	instr.Mnemonic = mnemonic
}

// Size returns the number of bytes consumed for the instruction.
func (instr *Base) Size() byte {
	return instr.size
}

// Exec writes the instruction to the final assembler.
func (instr *Base) Exec(a *asm.Assembler) {
	start := a.Len()

	switch instr.Mnemonic {
	case mnemonics.RET:
		a.Return()

	case mnemonics.SYSCALL:
		a.Syscall()

	case mnemonics.CPUID:
		a.CPUID()
	}

	instr.size = byte(a.Len() - start)
}

// String implements the string serialization.
func (instr *Base) String() string {
	return fmt.Sprintf("[%d]   %s", instr.size, mnemonicColor.Sprint(instr.Mnemonic))
}

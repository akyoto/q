package instructions

import (
	"fmt"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/assembler/mnemonics"
	"github.com/akyoto/q/build/register"
)

// MemoryRegister is used for instructions requiring a memory and a number operand.
type MemoryRegister struct {
	Base
	Destination *register.Register
	Source      *register.Register
	UsedBy1     string
	UsedBy2     string
	Offset      byte
	ByteCount   byte
}

// Exec writes the instruction to the final assembler.
func (instr *MemoryRegister) Exec(a *asm.Assembler) {
	start := a.Position()

	switch instr.Mnemonic {
	case mnemonics.STORE:
		a.StoreRegister(instr.Destination.Name, instr.Offset, instr.ByteCount, instr.Source.Name)

	default:
		panic("This should never happen!")
	}

	instr.size = byte(a.Position() - start)
}

// String implements the string serialization.
func (instr *MemoryRegister) String() string {
	return fmt.Sprintf("%s %dB [%v+%d], %s", mnemonicColor.Sprint(instr.Mnemonic), instr.ByteCount, instr.Destination.StringWithUser(instr.UsedBy1), instr.Offset, instr.Source.StringWithUser(instr.UsedBy2))
}

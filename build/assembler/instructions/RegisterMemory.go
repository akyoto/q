package instructions

import (
	"fmt"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/assembler/mnemonics"
	"github.com/akyoto/q/build/register"
)

// RegisterMemory is used for instructions requiring a register and a number operand.
type RegisterMemory struct {
	Base
	Destination *register.Register
	Source      *register.Register
	UsedBy1     string
	UsedBy2     string
	Offset      byte
	ByteCount   byte
}

// Exec writes the instruction to the final assembler.
func (instr *RegisterMemory) Exec(a *asm.Assembler) {
	start := a.Len()

	switch instr.Mnemonic {
	case mnemonics.LOAD:
		a.LoadRegister(instr.Destination.Name, instr.Source.Name, instr.Offset, instr.ByteCount)

	default:
		panic("This should never happen!")
	}

	instr.size = byte(a.Len() - start)
}

// String implements the string serialization.
func (instr *RegisterMemory) String() string {
	return fmt.Sprintf("[%d]   %s %dB %v, [%v+%d]", instr.size, mnemonicColor.Sprint(instr.Mnemonic), instr.ByteCount, instr.Destination.StringWithUser(instr.UsedBy1), instr.Source.StringWithUser(instr.UsedBy2), instr.Offset)
}

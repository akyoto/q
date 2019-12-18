package assembler

import (
	"fmt"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/register"
)

// registerMemory is used for instructions requiring a register and a number operand.
type registerMemory struct {
	Mnemonic    string
	Destination *register.Register
	Source      *register.Register
	UsedBy1     string
	UsedBy2     string
	Offset      byte
	ByteCount   byte
	size        byte
}

// Exec writes the instruction to the final assembler.
func (instr *registerMemory) Exec(a *asm.Assembler) {
	start := a.Len()

	switch instr.Mnemonic {
	case LOAD:
		a.LoadRegister(instr.Destination.Name, instr.Source.Name, instr.Offset, instr.ByteCount)

	default:
		panic("This should never happen!")
	}

	instr.size = byte(a.Len() - start)
}

// Name returns the mnemonic.
func (instr *registerMemory) Name() string {
	return instr.Mnemonic
}

// SetName sets the mnemonic.
func (instr *registerMemory) SetName(mnemonic string) {
	instr.Mnemonic = mnemonic
}

// Size returns the number of bytes consumed for the instruction.
func (instr *registerMemory) Size() byte {
	return instr.size
}

// String implements the string serialization.
func (instr *registerMemory) String() string {
	return fmt.Sprintf("[%d]   %s %dB %v, [%v+%d]", instr.size, mnemonicColor.Sprint(instr.Mnemonic), instr.ByteCount, instr.Destination.StringWithUser(instr.UsedBy1), instr.Source.StringWithUser(instr.UsedBy2), instr.Offset)
}

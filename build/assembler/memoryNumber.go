package assembler

import (
	"fmt"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/register"
)

// memoryNumber is used for instructions requiring a memory and a number operand.
type memoryNumber struct {
	Mnemonic    string
	Destination *register.Register
	Number      uint64
	UsedBy      string
	ByteCount   byte
	size        byte
}

// Exec writes the instruction to the final assembler.
func (instr *memoryNumber) Exec(a *asm.Assembler) {
	start := a.Len()

	switch instr.Mnemonic {
	case MOV:
		a.MoveMemoryNumber(instr.Destination.Name, instr.ByteCount, instr.Number)

	default:
		panic("This should never happen!")
	}

	instr.size = byte(a.Len() - start)
}

// Name returns the mnemonic.
func (instr *memoryNumber) Name() string {
	return instr.Mnemonic
}

// SetName sets the mnemonic.
func (instr *memoryNumber) SetName(mnemonic string) {
	instr.Mnemonic = mnemonic
}

// Size returns the number of bytes consumed for the instruction.
func (instr *memoryNumber) Size() byte {
	return instr.size
}

// String implements the string serialization.
func (instr *memoryNumber) String() string {
	return fmt.Sprintf("[%d]   %s [%v], %d", instr.size, mnemonicColor.Sprint(instr.Mnemonic), instr.Destination.StringWithUser(instr.UsedBy), instr.Number)
}

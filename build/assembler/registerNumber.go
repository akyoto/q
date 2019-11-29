package assembler

import (
	"fmt"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/register"
)

// registerNumber is used for instructions requiring a register and a number operand.
type registerNumber struct {
	Mnemonic    string
	Destination *register.Register
	Number      uint64
	UsedBy      fmt.Stringer
	size        byte
}

// Exec writes the instruction to the final assembler.
func (instr *registerNumber) Exec(a *asm.Assembler) {
	start := a.Len()

	switch instr.Mnemonic {
	case MOV:
		a.MoveRegisterNumber(instr.Destination.Name, instr.Number)

	case CMP:
		a.CompareRegisterNumber(instr.Destination.Name, instr.Number)

	case ADD:
		a.AddRegisterNumber(instr.Destination.Name, instr.Number)

	case MUL:
		a.MulRegisterNumber(instr.Destination.Name, instr.Number)

	case SUB:
		a.SubRegisterNumber(instr.Destination.Name, instr.Number)
	}

	instr.size = byte(a.Len() - start)
}

// Name returns the mnemonic.
func (instr *registerNumber) Name() string {
	return instr.Mnemonic
}

// Size returns the number of bytes consumed for the instruction.
func (instr *registerNumber) Size() byte {
	return instr.size
}

// String implements the string serialization.
func (instr *registerNumber) String() string {
	instr.Destination.ForceUse(instr.UsedBy)
	return fmt.Sprintf("[%d] %s %v, %d", instr.size, instr.Mnemonic, instr.Destination, instr.Number)
}

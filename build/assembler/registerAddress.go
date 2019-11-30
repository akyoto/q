package assembler

import (
	"fmt"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/register"
)

// registerAddress is used for instructions requiring Address register operands.
type registerAddress struct {
	Mnemonic    string
	Destination *register.Register
	UsedBy      string
	Address     uint32
	size        byte
}

// Exec writes the instruction to the final assembler.
func (instr *registerAddress) Exec(a *asm.Assembler) {
	start := a.Len()

	//nolint:gocritic
	switch instr.Mnemonic {
	case MOV:
		a.MoveRegisterAddress(instr.Destination.Name, instr.Address)
	}

	instr.size = byte(a.Len() - start)
}

// Name returns the mnemonic.
func (instr *registerAddress) Name() string {
	return instr.Mnemonic
}

// Size returns the number of bytes consumed for the instruction.
func (instr *registerAddress) Size() byte {
	return instr.size
}

// String implements the string serialization.
func (instr *registerAddress) String() string {
	return fmt.Sprintf("[%d] %s %v, <%v>", instr.size, instr.Mnemonic, instr.Destination.StringWithUser(instr.UsedBy), instr.Address)
}

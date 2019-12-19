package instructions

import (
	"fmt"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/assembler/mnemonics"
	"github.com/akyoto/q/build/register"
)

// RegisterAddress is used for instructions requiring Address register operands.
type RegisterAddress struct {
	Base
	Destination *register.Register
	UsedBy      string
	Address     uint32
}

// Exec writes the instruction to the final assembler.
func (instr *RegisterAddress) Exec(a *asm.Assembler) {
	start := a.Len()

	//nolint:gocritic
	switch instr.Mnemonic {
	case mnemonics.MOV:
		a.MoveRegisterAddress(instr.Destination.Name, instr.Address)
	}

	instr.size = byte(a.Len() - start)
}

// String implements the string serialization.
func (instr *RegisterAddress) String() string {
	return fmt.Sprintf("[%d]   %s %v, <%v>", instr.size, mnemonicColor.Sprint(instr.Mnemonic), instr.Destination.StringWithUser(instr.UsedBy), instr.Address)
}

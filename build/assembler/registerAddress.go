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
	Address     uint32
	cached      string
}

// Exec writes the instruction to the final assembler.
func (instr *registerAddress) Exec(a *asm.Assembler) {
	//nolint:gocritic
	switch instr.Mnemonic {
	case MOV:
		a.MoveRegisterAddress(instr.Destination.Name, instr.Address)
	}
}

// Name returns the mnemonic.
func (instr *registerAddress) Name() string {
	return instr.Mnemonic
}

// String implements the string serialization.
func (instr *registerAddress) String() string {
	if instr.cached != "" {
		return instr.cached
	}

	return fmt.Sprintf("%s %v, <%v>", instr.Mnemonic, instr.Destination, instr.Address)
}

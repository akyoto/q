package assembler

import (
	"fmt"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/register"
)

// register1 is used for instructions requiring 1 register operand.
type register1 struct {
	Mnemonic    string
	Destination *register.Register
	cached      string
}

// Exec writes the instruction to the final assembler.
func (instr *register1) Exec(a *asm.Assembler) {
	switch instr.Mnemonic {
	case INC:
		a.IncreaseRegister(instr.Destination.Name)

	case DEC:
		a.DecreaseRegister(instr.Destination.Name)

	case PUSH:
		a.PushRegister(instr.Destination.Name)

	case POP:
		a.PopRegister(instr.Destination.Name)
	}
}

// Name returns the mnemonic.
func (instr *register1) Name() string {
	return instr.Mnemonic
}

// String implements the string serialization.
func (instr *register1) String() string {
	if instr.cached != "" {
		return instr.cached
	}

	return fmt.Sprintf("%s %v", instr.Mnemonic, instr.Destination)
}

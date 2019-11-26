package assembler

import (
	"fmt"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/register"
)

// register2 is used for instructions requiring 2 register operands.
type register2 struct {
	Mnemonic    string
	Destination *register.Register
	Source      *register.Register
	cached      string
}

// Exec writes the instruction to the final assembler.
func (instr *register2) Exec(a *asm.Assembler) {
	switch instr.Mnemonic {
	case MOV:
		a.MoveRegisterRegister(instr.Destination.Name, instr.Source.Name)

	case CMP:
		a.CompareRegisterRegister(instr.Destination.Name, instr.Source.Name)

	case ADD:
		a.AddRegisterRegister(instr.Destination.Name, instr.Source.Name)

	case SUB:
		a.SubRegisterRegister(instr.Destination.Name, instr.Source.Name)

	case MUL:
		a.MulRegisterRegister(instr.Destination.Name, instr.Source.Name)
	}
}

// Name returns the mnemonic.
func (instr *register2) Name() string {
	return instr.Mnemonic
}

// String implements the string serialization.
func (instr *register2) String() string {
	if instr.cached != "" {
		return instr.cached
	}

	return fmt.Sprintf("%s %v, %v", instr.Mnemonic, instr.Destination, instr.Source)
}

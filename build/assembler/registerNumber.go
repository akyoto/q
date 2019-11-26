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
	cached      string
}

// Exec writes the instruction to the final assembler.
func (instr *registerNumber) Exec(a *asm.Assembler) {
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
}

// Name returns the mnemonic.
func (instr *registerNumber) Name() string {
	return instr.Mnemonic
}

// String implements the string serialization.
func (instr *registerNumber) String() string {
	if instr.cached != "" {
		return instr.cached
	}

	return fmt.Sprintf("%s %v, %d", instr.Mnemonic, instr.Destination, instr.Number)
}

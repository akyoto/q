package instructions

import (
	"fmt"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/assembler/mnemonics"
	"github.com/akyoto/q/build/register"
)

// RegisterNumber is used for instructions requiring a register and a number operand.
type RegisterNumber struct {
	Base
	Destination *register.Register
	Number      uint64
	UsedBy      string
}

// Exec writes the instruction to the final assembler.
func (instr *RegisterNumber) Exec(a *asm.Assembler) {
	start := a.Len()

	switch instr.Mnemonic {
	case mnemonics.MOV:
		a.MoveRegisterNumber(instr.Destination.Name, instr.Number)

	case mnemonics.CMP:
		a.CompareRegisterNumber(instr.Destination.Name, instr.Number)

	case mnemonics.ADD:
		a.AddRegisterNumber(instr.Destination.Name, instr.Number)

	case mnemonics.MUL:
		a.MulRegisterNumber(instr.Destination.Name, instr.Number)

	case mnemonics.SUB:
		a.SubRegisterNumber(instr.Destination.Name, instr.Number)
	}

	instr.size = byte(a.Len() - start)
}

// String implements the string serialization.
func (instr *RegisterNumber) String() string {
	return fmt.Sprintf("[%d]   %s %v, %d", instr.size, mnemonicColor.Sprint(instr.Mnemonic), instr.Destination.StringWithUser(instr.UsedBy), instr.Number)
}

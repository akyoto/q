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
	UsedBy      string
	size        byte
}

// Exec writes the instruction to the final assembler.
func (instr *register1) Exec(a *asm.Assembler) {
	start := a.Len()

	switch instr.Mnemonic {
	case INC:
		a.IncreaseRegister(instr.Destination.Name)

	case DEC:
		a.DecreaseRegister(instr.Destination.Name)

	case DIV:
		a.DivRegister(instr.Destination.Name)

	case CDQ:
		a.SignExtendToDX(instr.Destination.Name)

	case PUSH:
		a.PushRegister(instr.Destination.Name)

	case POP:
		a.PopRegister(instr.Destination.Name)
	}

	instr.size = byte(a.Len() - start)
}

// Name returns the mnemonic.
func (instr *register1) Name() string {
	return instr.Mnemonic
}

// Size returns the number of bytes consumed for the instruction.
func (instr *register1) Size() byte {
	return instr.size
}

// String implements the string serialization.
func (instr *register1) String() string {
	return fmt.Sprintf("[%d]   %s %v", instr.size, mnemonicColor.Sprint(instr.Mnemonic), instr.Destination.StringWithUser(instr.UsedBy))
}

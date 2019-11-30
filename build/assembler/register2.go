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
	UsedBy1     string
	UsedBy2     string
	size        byte
}

// Exec writes the instruction to the final assembler.
func (instr *register2) Exec(a *asm.Assembler) {
	start := a.Len()

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

	instr.size = byte(a.Len() - start)
}

// Name returns the mnemonic.
func (instr *register2) Name() string {
	return instr.Mnemonic
}

// Size returns the number of bytes consumed for the instruction.
func (instr *register2) Size() byte {
	return instr.size
}

// String implements the string serialization.
func (instr *register2) String() string {
	return fmt.Sprintf("[%d] %s %v, %v", instr.size, instr.Mnemonic, instr.Destination.StringWithUser(instr.UsedBy1), instr.Source.StringWithUser(instr.UsedBy2))
}

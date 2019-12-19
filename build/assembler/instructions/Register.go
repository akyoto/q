package instructions

import (
	"fmt"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/assembler/mnemonics"
	"github.com/akyoto/q/build/register"
)

// Register is used for instructions requiring 1 register operand.
type Register struct {
	Base
	Destination *register.Register
	UsedBy      string
}

// Exec writes the instruction to the final assembler.
func (instr *Register) Exec(a *asm.Assembler) {
	start := a.Len()

	switch instr.Mnemonic {
	case mnemonics.INC:
		a.IncreaseRegister(instr.Destination.Name)

	case mnemonics.DEC:
		a.DecreaseRegister(instr.Destination.Name)

	case mnemonics.DIV:
		a.DivRegister(instr.Destination.Name)

	case mnemonics.CDQ:
		a.SignExtendToDX(instr.Destination.Name)

	case mnemonics.PUSH:
		a.PushRegister(instr.Destination.Name)

	case mnemonics.POP:
		a.PopRegister(instr.Destination.Name)
	}

	instr.size = byte(a.Len() - start)
}

// String implements the string serialization.
func (instr *Register) String() string {
	return fmt.Sprintf("[%d]   %s %v", instr.size, mnemonicColor.Sprint(instr.Mnemonic), instr.Destination.StringWithUser(instr.UsedBy))
}

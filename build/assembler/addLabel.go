package assembler

import (
	"fmt"

	"github.com/akyoto/asm"
)

// addLabel is used for instructions that add a addLabel.
type addLabel struct {
	Label string
}

// Exec writes the instruction to the final assembler.
func (instr *addLabel) Exec(a *asm.Assembler) {
	a.AddLabel(instr.Label)
}

// Name returns the mnemonic.
func (instr *addLabel) Name() string {
	return ""
}

// Size returns the number of bytes consumed for the instruction.
func (instr *addLabel) Size() byte {
	return 0
}

// String implements the string serialization.
func (instr *addLabel) String() string {
	return fmt.Sprintf("[0] %s:", instr.Label)
}

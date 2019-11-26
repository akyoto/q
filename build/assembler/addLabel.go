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

// String implements the string serialization.
func (instr *addLabel) String() string {
	return fmt.Sprintf("%s:", instr.Label)
}

package assembler

import (
	"fmt"

	"github.com/akyoto/asm"
	"github.com/akyoto/color"
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

// SetName sets the mnemonic.
func (instr *addLabel) SetName(mnemonic string) {
	// Not applicable.
}

// Size returns the number of bytes consumed for the instruction.
func (instr *addLabel) Size() byte {
	return 0
}

// String implements the string serialization.
func (instr *addLabel) String() string {
	faint := color.New(color.Faint)
	return fmt.Sprintf("[0] %s:", faint.Sprint(instr.Label))
}

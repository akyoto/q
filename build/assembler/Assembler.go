package assembler

import (
	"log"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/register"
)

// Assembler produces machine code.
type Assembler struct {
	instructions      []instruction
	final             *asm.Assembler
	usedRegisterNames map[string]struct{}
	verbose           bool
}

// New creates a new assembler.
func New(verbose bool) *Assembler {
	return &Assembler{
		instructions:      make([]instruction, 0, 8),
		final:             asm.New(),
		usedRegisterNames: make(map[string]struct{}),
		verbose:           verbose,
	}
}

// AddLabel adds an instruction that adds a label.
func (a *Assembler) AddLabel(labelName string) {
	a.instructions = append(a.instructions, &addLabel{labelName})
}

// AddString adds a string.
func (a *Assembler) AddString(text string) uint32 {
	return a.final.Strings.Add(text)
}

// Finalize generates the final assembly code.
func (a *Assembler) Finalize() *asm.Assembler {
	for _, instr := range a.instructions {
		instr.Exec(a.final)
	}

	return a.final
}

// UsedRegisterNames returns the names of used registers.
func (a *Assembler) UsedRegisterNames() map[string]struct{} {
	return a.usedRegisterNames
}

// WriteTo generates the final assembly code.
func (a *Assembler) WriteTo(logger *log.Logger) {
	for _, instr := range a.instructions {
		logger.Println(instr.String())
	}
}

// lastInstruction returns the last added instruction.
func (a *Assembler) lastInstruction() instruction {
	if len(a.instructions) == 0 {
		return nil
	}

	return a.instructions[len(a.instructions)-1]
}

// removeLastInstruction removes the last added instruction.
func (a *Assembler) removeLastInstruction() {
	if len(a.instructions) == 0 {
		return
	}

	a.instructions = a.instructions[:len(a.instructions)-1]
}

// do adds an instruction without any operands.
func (a *Assembler) do(mnemonic string) {
	a.instructions = append(a.instructions, &standalone{mnemonic, 0})
}

// doRegister1 adds an instruction with a single register operand.
func (a *Assembler) doRegister1(mnemonic string, destination *register.Register) {
	instr := &register1{
		Mnemonic:    mnemonic,
		Destination: destination,
	}

	if a.verbose {
		instr.UsedBy = destination.UserString()
	}

	a.instructions = append(a.instructions, instr)
	a.usedRegisterNames[destination.Name] = nothing
}

// doRegister2 adds an instruction using 2 registers.
func (a *Assembler) doRegister2(mnemonic string, destination *register.Register, source *register.Register) {
	instr := &register2{
		Mnemonic:    mnemonic,
		Destination: destination,
		Source:      source,
	}

	if a.verbose {
		instr.UsedBy1 = destination.UserString()
		instr.UsedBy2 = source.UserString()
	}

	a.instructions = append(a.instructions, instr)
	a.usedRegisterNames[destination.Name] = nothing
	a.usedRegisterNames[source.Name] = nothing
}

// doRegisterNumber adds an instruction using a register and a number.
func (a *Assembler) doRegisterNumber(mnemonic string, destination *register.Register, number uint64) {
	instr := &registerNumber{
		Mnemonic:    mnemonic,
		Destination: destination,
		Number:      number,
	}

	if a.verbose {
		instr.UsedBy = destination.UserString()
	}

	a.instructions = append(a.instructions, instr)
	a.usedRegisterNames[destination.Name] = nothing
}

// doRegisterAddress adds an instruction using a register and a section address.
func (a *Assembler) doRegisterAddress(mnemonic string, destination *register.Register, address uint32) {
	instr := &registerAddress{
		Mnemonic:    mnemonic,
		Destination: destination,
		Address:     address,
	}

	if a.verbose {
		instr.UsedBy = destination.UserString()
	}

	a.instructions = append(a.instructions, instr)
	a.usedRegisterNames[destination.Name] = nothing
}

// doLabel adds an instruction with a label operand.
func (a *Assembler) doLabel(mnemonic string, labelName string) {
	a.instructions = append(a.instructions, &label{mnemonic, labelName, 0})
}

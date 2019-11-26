package assembler

import (
	"log"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/register"
)

// Assembler produces machine code.
type Assembler struct {
	instructions []instruction
	final        *asm.Assembler
	logger       *log.Logger
}

func New(logger *log.Logger) *Assembler {
	return &Assembler{
		final:  asm.New(),
		logger: logger,
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
func (a *Assembler) Finalize(logger *log.Logger) *asm.Assembler {
	for _, instr := range a.instructions {
		instr.Exec(a.final)

		if logger != nil {
			logger.Println(instr.String())
		}
	}

	return a.final
}

// lastInstruction returns the last added instruction.
func (a *Assembler) lastInstruction() instruction {
	if len(a.instructions) == 0 {
		return nil
	}

	return a.instructions[len(a.instructions)-1]
}

// do adds an instruction without any operands.
func (a *Assembler) do(mnemonic string) {
	a.instructions = append(a.instructions, &standalone{mnemonic})
}

// doRegister1 adds an instruction with a single register operand.
func (a *Assembler) doRegister1(mnemonic string, destination *register.Register) {
	instr := &register1{
		Mnemonic:    mnemonic,
		Destination: destination,
	}

	if a.logger != nil {
		instr.cached = instr.String()
	}

	a.instructions = append(a.instructions, instr)
}

// doRegister2 adds an instruction using 2 registers.
func (a *Assembler) doRegister2(mnemonic string, destination *register.Register, source *register.Register) {
	instr := &register2{
		Mnemonic:    mnemonic,
		Destination: destination,
		Source:      source,
	}

	if a.logger != nil {
		instr.cached = instr.String()
	}

	a.instructions = append(a.instructions, instr)
}

// doRegisterNumber adds an instruction using a register and a number.
func (a *Assembler) doRegisterNumber(mnemonic string, destination *register.Register, number uint64) {
	instr := &registerNumber{
		Mnemonic:    mnemonic,
		Destination: destination,
		Number:      number,
	}

	if a.logger != nil {
		instr.cached = instr.String()
	}

	a.instructions = append(a.instructions, instr)
}

// doRegisterAddress adds an instruction using a register and a section address.
func (a *Assembler) doRegisterAddress(mnemonic string, destination *register.Register, address uint32) {
	instr := &registerAddress{
		Mnemonic:    mnemonic,
		Destination: destination,
		Address:     address,
	}

	if a.logger != nil {
		instr.cached = instr.String()
	}

	a.instructions = append(a.instructions, instr)
}

// doLabel adds an instruction with a label operand.
func (a *Assembler) doLabel(mnemonic string, labelName string) {
	a.instructions = append(a.instructions, &label{mnemonic, labelName})
}

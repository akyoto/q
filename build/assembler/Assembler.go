package assembler

import (
	"log"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/register"
)

// Assembler produces machine code.
type Assembler struct {
	Instructions    []instruction
	final           *asm.Assembler
	usedRegisterIDs map[byte]struct{}
	verbose         bool
}

// New creates a new assembler.
func New(verbose bool) *Assembler {
	return &Assembler{
		Instructions:    make([]instruction, 0, 8),
		final:           asm.New(),
		usedRegisterIDs: make(map[byte]struct{}),
		verbose:         verbose,
	}
}

// AddLabel adds an instruction that adds a label.
func (a *Assembler) AddLabel(labelName string) {
	jump, isJump := a.lastInstruction().(*label)

	if isJump && jump.Label == labelName {
		a.removeLastInstruction()
	}

	a.Instructions = append(a.Instructions, &addLabel{labelName})
}

// AddString adds a string.
func (a *Assembler) AddString(text string) uint32 {
	return a.final.Strings.Add(text)
}

// Finalize generates the final assembly code.
func (a *Assembler) Finalize() *asm.Assembler {
	for _, instr := range a.Instructions {
		instr.Exec(a.final)
	}

	return a.final
}

// UsedRegisterIDs returns the IDs of used registers.
func (a *Assembler) UsedRegisterIDs() map[byte]struct{} {
	return a.usedRegisterIDs
}

// WriteTo generates the final assembly code.
func (a *Assembler) WriteTo(logger *log.Logger) {
	for _, instr := range a.Instructions {
		logger.Println(instr.String())
	}
}

// lastInstruction returns the last added instruction.
func (a *Assembler) lastInstruction() instruction {
	if len(a.Instructions) == 0 {
		return nil
	}

	return a.Instructions[len(a.Instructions)-1]
}

// removeLastInstruction removes the last added instruction.
func (a *Assembler) removeLastInstruction() {
	if len(a.Instructions) == 0 {
		return
	}

	a.Instructions = a.Instructions[:len(a.Instructions)-1]
}

// do adds an instruction without any operands.
func (a *Assembler) do(mnemonic string) {
	a.Instructions = append(a.Instructions, &standalone{mnemonic, 0})
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

	a.Instructions = append(a.Instructions, instr)
	a.usedRegisterIDs[destination.ID] = nothing
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

	a.Instructions = append(a.Instructions, instr)
	a.usedRegisterIDs[destination.ID] = nothing
	a.usedRegisterIDs[source.ID] = nothing
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

	a.Instructions = append(a.Instructions, instr)
	a.usedRegisterIDs[destination.ID] = nothing
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

	a.Instructions = append(a.Instructions, instr)
	a.usedRegisterIDs[destination.ID] = nothing
}

// doMemoryNumber adds an instruction using a memory address and a number.
func (a *Assembler) doMemoryNumber(mnemonic string, destination *register.Register, offset byte, byteCount byte, number uint64) {
	instr := &memoryNumber{
		Mnemonic:    mnemonic,
		Destination: destination,
		Offset:      offset,
		ByteCount:   byteCount,
		Number:      number,
	}

	if a.verbose {
		instr.UsedBy = destination.UserString()
	}

	a.Instructions = append(a.Instructions, instr)
	a.usedRegisterIDs[destination.ID] = nothing
}

// doLabel adds an instruction with a label operand.
func (a *Assembler) doLabel(mnemonic string, labelName string) {
	a.Instructions = append(a.Instructions, &label{mnemonic, labelName, 0})
}

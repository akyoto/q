package assembler

import (
	"log"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/assembler/instructions"
	"github.com/akyoto/q/build/register"
)

// Assembler produces machine code.
type Assembler struct {
	Instructions    []instruction
	usedRegisterIDs []register.ID
	final           *asm.Assembler
	verbose         bool
}

// New creates a new assembler.
func New(verbose bool) *Assembler {
	return &Assembler{
		Instructions: make([]instruction, 0, 8),
		final:        asm.New(),
		verbose:      verbose,
	}
}

// AddLabel adds an instruction that adds a label.
func (a *Assembler) AddLabel(labelName string) {
	jump, isJump := a.lastInstruction().(*instructions.Jump)

	if isJump && jump.Label == labelName {
		a.removeLastInstruction()
	}

	a.Instructions = append(a.Instructions, &instructions.AddLabel{Label: labelName})
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

// UseRegisterID marks the given register ID as used.
func (a *Assembler) UseRegisterID(newID register.ID) {
	for _, id := range a.usedRegisterIDs {
		if id == newID {
			return
		}
	}

	a.usedRegisterIDs = append(a.usedRegisterIDs, newID)
}

// UsedRegisterIDs returns the IDs of used registers.
func (a *Assembler) UsedRegisterIDs() []register.ID {
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
	a.Instructions = append(a.Instructions, &instructions.Base{Mnemonic: mnemonic})
}

// doRegister adds an instruction with a single register operand.
func (a *Assembler) doRegister(mnemonic string, destination *register.Register) {
	instr := &instructions.Register{
		Destination: destination,
	}

	instr.SetName(mnemonic)

	if a.verbose {
		instr.UsedBy = destination.UserString()
	}

	a.Instructions = append(a.Instructions, instr)
	a.UseRegisterID(destination.ID)
}

// doRegisterRegister adds an instruction using 2 registers.
func (a *Assembler) doRegisterRegister(mnemonic string, destination *register.Register, source *register.Register) {
	instr := &instructions.RegisterRegister{
		Destination: destination,
		Source:      source,
	}

	instr.SetName(mnemonic)

	if a.verbose {
		instr.UsedBy1 = destination.UserString()
		instr.UsedBy2 = source.UserString()
	}

	a.Instructions = append(a.Instructions, instr)
	a.UseRegisterID(destination.ID)
	a.UseRegisterID(source.ID)
}

// doRegisterNumber adds an instruction using a register and a number.
func (a *Assembler) doRegisterNumber(mnemonic string, destination *register.Register, number uint64) {
	instr := &instructions.RegisterNumber{
		Destination: destination,
		Number:      number,
	}

	instr.SetName(mnemonic)

	if a.verbose {
		instr.UsedBy = destination.UserString()
	}

	a.Instructions = append(a.Instructions, instr)
	a.UseRegisterID(destination.ID)
}

// doRegisterAddress adds an instruction using a register and a section address.
func (a *Assembler) doRegisterAddress(mnemonic string, destination *register.Register, address uint32) {
	instr := &instructions.RegisterAddress{
		Destination: destination,
		Address:     address,
	}

	instr.SetName(mnemonic)

	if a.verbose {
		instr.UsedBy = destination.UserString()
	}

	a.Instructions = append(a.Instructions, instr)
	a.UseRegisterID(destination.ID)
}

// doMemoryNumber adds an instruction using a memory address and a number.
func (a *Assembler) doMemoryNumber(mnemonic string, destination *register.Register, offset byte, byteCount byte, number uint64) {
	instr := &instructions.MemoryNumber{
		Destination: destination,
		Offset:      offset,
		ByteCount:   byteCount,
		Number:      number,
	}

	instr.SetName(mnemonic)

	if a.verbose {
		instr.UsedBy = destination.UserString()
	}

	a.Instructions = append(a.Instructions, instr)
	a.UseRegisterID(destination.ID)
}

// doMemoryRegister adds an instruction using a memory address and a register.
func (a *Assembler) doMemoryRegister(mnemonic string, destination *register.Register, offset byte, byteCount byte, source *register.Register) {
	instr := &instructions.MemoryRegister{
		Destination: destination,
		Offset:      offset,
		ByteCount:   byteCount,
		Source:      source,
	}

	instr.SetName(mnemonic)

	if a.verbose {
		instr.UsedBy1 = destination.UserString()
		instr.UsedBy2 = source.UserString()
	}

	a.Instructions = append(a.Instructions, instr)
	a.UseRegisterID(destination.ID)
	a.UseRegisterID(source.ID)
}

// doRegisterMemory adds an instruction using a register target and a source memory address.
func (a *Assembler) doRegisterMemory(mnemonic string, destination *register.Register, source *register.Register, offset byte, byteCount byte) {
	instr := &instructions.RegisterMemory{
		Destination: destination,
		Source:      source,
		Offset:      offset,
		ByteCount:   byteCount,
	}

	instr.SetName(mnemonic)

	if a.verbose {
		instr.UsedBy1 = destination.UserString()
		instr.UsedBy2 = source.UserString()
	}

	a.Instructions = append(a.Instructions, instr)
	a.UseRegisterID(destination.ID)
	a.UseRegisterID(source.ID)
}

// doJump adds a jump instruction with a label operand.
func (a *Assembler) doJump(mnemonic string, labelName string) {
	instr := &instructions.Jump{Label: labelName}
	instr.SetName(mnemonic)

	a.Instructions = append(a.Instructions, instr)
}

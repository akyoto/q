package asm

import (
	"maps"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/data"
	"git.urbach.dev/cli/q/src/elf"
	"git.urbach.dev/cli/q/src/exe"
)

// Assembler contains a list of instructions.
type Assembler struct {
	Data         data.Data
	Instructions []Instruction
}

// Append adds another instruction.
func (a *Assembler) Append(instr Instruction) {
	if a.Skip(instr) {
		return
	}

	a.Instructions = append(a.Instructions, instr)
}

// Last returns the last instruction.
func (a *Assembler) Last() Instruction {
	return a.Instructions[len(a.Instructions)-1]
}

// Compile compiles the instructions to machine code.
func (a *Assembler) Compile(b *build.Build) (code []byte, data []byte) {
	data, dataLabels := a.Data.Finalize()

	c := compiler{
		code:       make([]byte, 0, len(a.Instructions)*8),
		data:       data,
		dataLabels: dataLabels,
		labels:     make(map[string]Address, 32),
	}

	switch b.Arch {
	case build.ARM:
		armc := compilerARM{compiler: &c}

		for _, instr := range a.Instructions {
			armc.Compile(instr)
		}

	case build.X86:
		x86c := compilerX86{compiler: &c}

		for _, instr := range a.Instructions {
			x86c.Compile(instr)
		}
	}

	x := exe.New(elf.HeaderEnd, b.FileAlign, b.MemoryAlign)
	x.InitSections(c.code, c.data)
	dataSectionOffset := x.Sections[1].MemoryOffset - x.Sections[0].MemoryOffset

	for dataLabel, address := range dataLabels {
		c.labels[dataLabel] = dataSectionOffset + address
	}

	for _, call := range c.deferred {
		call()
	}

	return c.code, c.data
}

// Merge combines the contents of this assembler with another one.
func (a *Assembler) Merge(b *Assembler) {
	maps.Copy(a.Data, b.Data)

	for _, instr := range b.Instructions {
		a.Append(instr)
	}
}

// SetData sets the data for the given label.
func (a *Assembler) SetData(label string, bytes []byte) {
	if a.Data == nil {
		a.Data = data.Data{}
	}

	a.Data.Insert(label, bytes)
}

// SetLast sets the last instruction.
func (a *Assembler) SetLast(instr Instruction) {
	a.Instructions[len(a.Instructions)-1] = instr
}

// Skip returns true if appending the instruction can be skipped.
func (a *Assembler) Skip(instr Instruction) bool {
	if len(a.Instructions) == 0 {
		return false
	}

	switch instr := instr.(type) {
	case *FunctionEnd:
		// Call + FunctionEnd can be replaced by a single Jump
		call, isCall := a.Last().(*Call)

		if isCall {
			a.SetLast(&Jump{Label: call.Label})
			return true
		}

	case *Label:
		// Jump + Label can be replaced by just the Label if both addresses are equal
		jump, isJump := a.Last().(*Jump)

		if isJump && jump.Label == instr.Name {
			a.SetLast(instr)
			return true
		}

	case *Return:
		// Call + Return can be replaced by a single Jump
		call, isCall := a.Last().(*Call)

		if isCall {
			a.SetLast(&Jump{Label: call.Label})
			return true
		}

		// Jump + Return is unnecessary
		_, isJump := a.Last().(*Jump)

		if isJump {
			return true
		}

		// Return + Return is unnecessary
		_, isReturn := a.Last().(*Return)

		if isReturn {
			return true
		}
	}

	return false
}
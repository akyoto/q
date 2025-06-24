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
	a.Instructions = append(a.Instructions, instr)
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
	a.Instructions = append(a.Instructions, b.Instructions...)
}

// SetData sets the data for the given label.
func (a *Assembler) SetData(label string, bytes []byte) {
	if a.Data == nil {
		a.Data = data.Data{}
	}

	a.Data.Insert(label, bytes)
}
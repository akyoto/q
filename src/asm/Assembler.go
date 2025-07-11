package asm

import (
	"maps"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/data"
	"git.urbach.dev/cli/q/src/dll"
)

// Assembler contains a list of instructions.
type Assembler struct {
	Data         data.Data
	Instructions []Instruction
	Libraries    dll.List
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
func (a *Assembler) Compile(b *build.Build) (code []byte, data []byte, libs dll.List) {
	data, dataLabels := a.Data.Finalize()

	c := compiler{
		patcher: patcher{
			code:         make([]byte, 0, len(a.Instructions)*8),
			earlyPatches: make([]*patch, 0, len(a.Instructions)/8),
			latePatches:  make([]*patch, 0, len(a.Instructions)/8),
			labels:       make(map[string]int, 32),
		},
		build:      b,
		data:       data,
		dataLabels: dataLabels,
		libraries:  a.Libraries,
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

	c.ApplyPatches(c.earlyPatches)
	c.AddDataLabels()
	c.ApplyPatches(c.latePatches)
	return c.code, c.data, c.libraries
}

// Merge combines the contents of this assembler with another one.
func (a *Assembler) Merge(b *Assembler) {
	skip := 0

	for a.Skip(b.Instructions[skip]) {
		skip++
	}

	a.Instructions = append(a.Instructions, b.Instructions[skip:]...)
	maps.Copy(a.Data, b.Data)

	for _, library := range b.Libraries {
		for _, fn := range library.Functions {
			a.Libraries = a.Libraries.Append(library.Name, fn)
		}
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
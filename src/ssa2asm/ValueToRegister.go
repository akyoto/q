package ssa2asm

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
)

// ValueToRegister moves a value into the given `destination` register.
func (f *Compiler) ValueToRegister(instr ssa.Value, destination cpu.Register) {
	switch instr := instr.(type) {
	case *ssa.Bytes:
		f.Count.Data++
		label := f.CreateLabel("data", f.Count.Data)
		f.Assembler.SetData(label.Name, instr.Bytes)

		f.Assembler.Append(&asm.MoveRegisterLabel{
			Destination: destination,
			Label:       label.Name,
		})

	case *ssa.Int:
		f.Assembler.Append(&asm.MoveRegisterNumber{
			Destination: destination,
			Number:      instr.Int,
		})

	case *ssa.Parameter:
		source := f.CPU.Call[instr.Index]

		if source == destination {
			return
		}

		f.Assembler.Append(&asm.MoveRegisterRegister{
			Destination: destination,
			Source:      source,
		})

	case *ssa.StructField:
		parameter := instr.Struct.(*ssa.Parameter)
		field := instr.Field
		source := f.CPU.Call[parameter.Index+field.Index]

		if source == destination {
			return
		}

		f.Assembler.Append(&asm.MoveRegisterRegister{
			Destination: destination,
			Source:      source,
		})

	case *ssa.Syscall:
		if destination == f.CPU.Return[0] {
			return
		}

		f.Assembler.Append(&asm.MoveRegisterRegister{
			Destination: destination,
			Source:      f.CPU.Return[0],
		})
	}
}
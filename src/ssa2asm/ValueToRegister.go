package ssa2asm

import (
	"slices"
	"strings"

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

	case *ssa.Call:
		fn := instr.Arguments[0].(*ssa.Function)
		args := instr.Arguments[1:]

		if fn.IsExtern {
			for i := range slices.Backward(args) {
				f.ValueToRegister(args[i], f.CPU.ExternCall[i])
			}

			dot := strings.IndexByte(fn.UniqueName, '.')
			library := fn.UniqueName[:dot]
			function := fn.UniqueName[dot+1:]
			f.Assembler.Append(&asm.CallExtern{Library: library, Function: function})
		} else {
			offset := 0

			for i := range slices.Backward(args) {
				structure, isStruct := args[i].(*ssa.Struct)

				if isStruct {
					for _, field := range structure.Arguments {
						f.ValueToRegister(field, f.CPU.Call[offset+i])
						i++
					}
				} else {
					f.ValueToRegister(args[i], f.CPU.Call[offset+i])
				}
			}

			f.Assembler.Append(&asm.Call{Label: fn.UniqueName})
		}

		if destination == f.CPU.Return[0] {
			return
		}

		f.Assembler.Append(&asm.MoveRegisterRegister{
			Destination: destination,
			Source:      f.CPU.Return[0],
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
		for i := range slices.Backward(instr.Arguments) {
			f.ValueToRegister(instr.Arguments[i], f.CPU.Syscall[i])
		}

		f.Assembler.Append(&asm.Syscall{})

		if destination == f.CPU.Return[0] {
			return
		}

		f.Assembler.Append(&asm.MoveRegisterRegister{
			Destination: destination,
			Source:      f.CPU.Return[0],
		})
	}
}
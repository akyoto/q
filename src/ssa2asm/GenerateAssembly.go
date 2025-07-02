package ssa2asm

import (
	"slices"
	"strings"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

// GenerateAssembly converts the SSA IR to assembler instructions.
func (f *Compiler) GenerateAssembly(ir ssa.IR, isLeaf bool) {
	f.Assembler.Append(&asm.Label{Name: f.UniqueName})

	if !isLeaf && f.UniqueName != "core.init" {
		f.Assembler.Append(&asm.FunctionStart{})
	}

	for instr := range ir.Values {
		switch instr := instr.(type) {
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

		case *ssa.Return:
			for i := range slices.Backward(instr.Arguments) {
				f.ValueToRegister(instr.Arguments[i], f.CPU.Return[i])
			}

			f.Assembler.Append(&asm.Return{})

		case *ssa.Syscall:
			for i := range slices.Backward(instr.Arguments) {
				f.ValueToRegister(instr.Arguments[i], f.CPU.Syscall[i])
			}

			f.Assembler.Append(&asm.Syscall{})
		}
	}

	if !isLeaf && f.UniqueName != "core.init" {
		f.Assembler.Append(&asm.FunctionEnd{})
	}

	if f.UniqueName != "core.exit" {
		f.Assembler.Append(&asm.Return{})
	}
}
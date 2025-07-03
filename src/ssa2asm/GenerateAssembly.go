package ssa2asm

import (
	"slices"

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
		if len(instr.Users()) > 0 {
			continue
		}

		switch instr := instr.(type) {
		case *ssa.Call, *ssa.CallExtern, *ssa.Syscall:
			f.ValueToRegister(instr, f.CPU.Return[0])

		case *ssa.Return:
			for i := range slices.Backward(instr.Arguments) {
				f.ValueToRegister(instr.Arguments[i], f.CPU.Return[i])
			}

			f.Assembler.Append(&asm.Return{})
		}
	}

	if !isLeaf && f.UniqueName != "core.init" {
		f.Assembler.Append(&asm.FunctionEnd{})
	}

	if f.UniqueName != "core.exit" {
		f.Assembler.Append(&asm.Return{})
	}
}
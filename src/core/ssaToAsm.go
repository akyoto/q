package core

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
)

// ssaToAsm converts the SSA IR to assembler instructions.
func (f *Function) ssaToAsm() {
	f.Assembler.Append(&asm.Label{Name: f.UniqueName})

	if !f.IsLeaf() && f.UniqueName != "core.init" {
		f.Assembler.Append(&asm.FunctionStart{})
	}

	for instr := range f.Values {
		switch instr := instr.(type) {
		case *ssa.Call:
			arg := instr.Arguments[0].(*ssa.Function)
			fn := f.All.Functions[arg.UniqueName]

			if fn.IsExtern() {
				f.ssaValuesToRegisters(instr.Arguments[1:], f.CPU.ExternCall)
				f.Assembler.Append(&asm.CallExtern{Library: fn.Package, Function: fn.Name})
			} else {
				f.ssaValuesToRegisters(instr.Arguments[1:], f.CPU.Call)
				f.Assembler.Append(&asm.Call{Label: fn.UniqueName})
			}

		case *ssa.Syscall:
			f.ssaValuesToRegisters(instr.Arguments, f.CPU.Syscall)
			f.Assembler.Append(&asm.Syscall{})

		case *ssa.Return:
			f.Assembler.Append(&asm.Return{})
		}
	}

	if !f.IsLeaf() && f.UniqueName != "core.init" {
		f.Assembler.Append(&asm.FunctionEnd{})
	}

	if f.UniqueName != "core.exit" {
		f.Assembler.Append(&asm.Return{})
	}
}
package core

import (
	"slices"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// Compile turns a function into machine code.
func (f *Function) Compile() {
	extra := 0

	for i, input := range f.Input {
		if input.Name == "_" {
			continue
		}

		f.Identifiers[input.Name] = f.AppendRegister(i + extra)

		if input.TypeTokens[0].Kind == token.ArrayStart {
			extra++
			f.Identifiers[input.Name+".length"] = f.AppendRegister(i + extra)
		}
	}

	for instr := range f.Body.Instructions {
		f.Err = f.CompileInstruction(instr)

		if f.Err != nil {
			return
		}
	}

	f.Err = f.CheckDeadCode()

	if f.Err != nil {
		return
	}

	f.Assembler.Append(&asm.Label{Name: f.UniqueName})

	if !f.IsLeaf() && f.UniqueName != "core.init" {
		f.Assembler.Append(&asm.FunctionStart{})
	}

	for instr := range f.Values {
		switch instr := instr.(type) {
		case *ssa.Call:
			f.mv(instr.Args[1:], f.CPU.Call)

			switch arg := instr.Args[0].(type) {
			case *ssa.Function:
				f.Assembler.Instructions = append(f.Assembler.Instructions, &asm.Call{Label: arg.UniqueName})
			}

		case *ssa.Syscall:
			f.mv(instr.Args, f.CPU.Syscall)
			f.Assembler.Append(&asm.Syscall{})

		case *ssa.Return:
			f.Assembler.Append(&asm.Return{})
		}
	}

	if !f.IsLeaf() && f.UniqueName != "core.init" {
		f.Assembler.Append(&asm.FunctionEnd{})
	}

	switch f.Assembler.Instructions[len(f.Assembler.Instructions)-1].(type) {
	case *asm.Return:
	default:
		f.Assembler.Append(&asm.Return{})
	}
}

func (f *Function) mv(args []ssa.Value, registers []cpu.Register) {
	extra := 0

	for _, arg := range args {
		switch arg.(type) {
		case *ssa.Bytes:
			extra++
		}
	}

	for i, arg := range slices.Backward(args) {
		switch arg := arg.(type) {
		case *ssa.Int:
			f.Assembler.Append(&asm.MoveRegisterNumber{
				Destination: registers[i+extra],
				Number:      arg.Int,
			})
		case *ssa.Parameter:
			f.Assembler.Append(&asm.MoveRegisterRegister{
				Destination: registers[i+extra],
				Source:      f.CPU.Call[arg.Index],
			})
		case *ssa.Bytes:
			f.Assembler.SetData("data0", arg.Bytes)

			f.Assembler.Append(&asm.MoveRegisterNumber{
				Destination: registers[i+extra],
				Number:      len(arg.Bytes),
			})

			extra--

			f.Assembler.Append(&asm.MoveRegisterLabel{
				Destination: registers[i+extra],
				Label:       "data0",
			})
		}
	}
}
package core

import (
	"slices"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
)

// ssaValuesToRegisters generates assembler instructions to move the SSA values to the given registers.
func (f *Function) ssaValuesToRegisters(args []ssa.Value, registers []cpu.Register) {
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
			f.count.data++
			label := f.CreateLabel("data", f.count.data)
			f.Assembler.SetData(label.Name, arg.Bytes)

			f.Assembler.Append(&asm.MoveRegisterNumber{
				Destination: registers[i+extra],
				Number:      len(arg.Bytes),
			})

			extra--

			f.Assembler.Append(&asm.MoveRegisterLabel{
				Destination: registers[i+extra],
				Label:       label.Name,
			})
		}
	}
}
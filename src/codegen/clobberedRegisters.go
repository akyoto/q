package codegen

import (
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// clobberedRegisters returns the registers that are clobbered by the given instruction.
func (f *Function) clobberedRegisters(instr ssa.Value) []cpu.Register {
	switch instr := instr.(type) {
	case *ssa.BinaryOp:
		switch instr.Op {
		case token.Div, token.Mod:
			return f.CPU.DivisionClobbered
		case token.Shl, token.Shr:
			return f.CPU.ShiftClobbered
		default:
			return nil
		}
	case *ssa.Call:
		return f.CPU.Call.Clobbered
	case *ssa.CallExtern:
		return f.CPU.ExternCall.Clobbered
	case *ssa.Cas:
		return f.CPU.CasClobbered
	case *ssa.Syscall:
		return append(f.CPU.Syscall.Clobbered, f.CPU.Syscall.In[:min(len(instr.Arguments), len(f.CPU.Syscall.In))]...)
	default:
		return nil
	}
}
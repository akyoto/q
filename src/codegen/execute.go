package codegen

import (
	"git.urbach.dev/cli/q/src/ssa"
)

// execute executes a step which appends it to the assembler's instruction list.
func (f *Function) execute(step *Step) {
	switch instr := step.Value.(type) {
	case *ssa.BinaryOp:
		f.executeBinaryOp(step, instr)
	case *ssa.Bool:
		f.executeBool(step, instr)
	case *ssa.Branch:
		f.executeBranch(step, instr)
	case *ssa.Bytes:
		f.executeBytes(step, instr)
	case *ssa.Call:
		f.executeCall(step, instr)
	case *ssa.CallExtern:
		f.executeCallExtern(step, instr)
	case *ssa.Copy:
		f.executeCopy(step, instr)
	case *ssa.Data:
		f.executeData(step, instr)
	case *ssa.FromTuple:
		f.executeFromTuple(step, instr)
	case *ssa.Function:
		f.executeFunction(step, instr)
	case *ssa.Int:
		f.executeInt(step, instr)
	case *ssa.Jump:
		f.executeJump(step, instr)
	case *Label:
		f.executeLabel(instr)
	case *ssa.Load:
		f.executeLoad(step, instr)
	case *ssa.Memory:
		f.executeMemory(step, instr)
	case *ssa.Parameter:
		f.executeParameter(step, instr)
	case *ssa.Phi:
		f.executePhi(instr)
	case *ssa.Register:
		f.executeRegister(step, instr)
	case *ssa.Return:
		f.executeReturn(instr)
	case *ssa.Store:
		f.executeStore(instr)
	case *ssa.Syscall:
		f.executeSyscall(step, instr)
	case *ssa.UnaryOp:
		f.executeUnaryOp(step, instr)
	default:
		panic("not implemented: " + instr.String())
	}
}
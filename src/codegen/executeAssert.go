package codegen

import (
	"git.urbach.dev/cli/q/src/ssa"
)

func (f *Function) executeAssert(instr *ssa.Assert) {
	f.jumpIfFalse(instr.Condition.(*ssa.BinaryOp).Op, "run.crash")
}
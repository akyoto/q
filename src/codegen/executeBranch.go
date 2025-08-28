package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

func (f *Function) executeBranch(step *Step, instr *ssa.Branch) {
	var op token.Kind
	binaryOp, isBinaryOp := instr.Condition.(*ssa.BinaryOp)

	if isBinaryOp && binaryOp.Op.IsComparison() {
		op = binaryOp.Op
	} else {
		op = token.NotEqual

		f.Assembler.Append(&asm.CompareNumber{
			Destination: f.ValueToStep[instr.Condition].Register,
			Number:      0,
		})
	}

	f.insertPhiMoves(step)
	following := f.Steps[step.Index+1].Value.(*Label)

	switch following.Name {
	case instr.Then.Label:
		f.jumpIfFalse(op, instr.Else.Label)
	case instr.Else.Label:
		f.jumpIfTrue(op, instr.Then.Label)
	default:
		panic("branch instruction must be followed by the 'then' or 'else' block")
	}
}
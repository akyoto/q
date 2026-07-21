package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
	"git.urbach.dev/cli/q/src/x86"
)

func (f *Function) executeBranch(step *Step, instr *ssa.Branch) {
	var (
		op            token.Kind
		unsigned      bool
		conditionStep = f.ValueToStep[instr.Condition]
	)

	if conditionStep.Register != -1 {
		op = token.NotEqual
		operand := f.resolveOperand(conditionStep, step.Live)

		f.Assembler.Append(&asm.CompareNumber{
			Destination: operand,
			Number:      0,
		})
	} else {
		switch condition := instr.Condition.(type) {
		case *ssa.BinaryOp:
			if condition.Op.IsComparison() {
				op = condition.Op
				unsigned = types.IsUnsigned(condition.Left.Type()) || types.IsUnsigned(condition.Right.Type())
			} else {
				panic("condition using a binary operation not assigned to a register")
			}

		case *ssa.Cas:
			op = token.Equal
			operand := f.ValueToStep[condition.Arguments[1]].Register

			if f.build.Arch == config.X86 {
				operand = x86.R0
			}

			f.Assembler.Append(&asm.CompareNumber{
				Destination: operand,
				Number:      condition.Arguments[1].(*ssa.Int).Int,
			})
		}
	}

	f.insertPhiMoves(step)
	following := f.Steps[step.Index+1].Value.(*Label)
	condition := tokenToCondition(op, unsigned)

	switch following.Name {
	case instr.Then.Label:
		f.jumpIfFalse(condition, instr.Else.Label)
	case instr.Else.Label:
		f.jumpIfTrue(condition, instr.Then.Label)
	default:
		panic("branch instruction must be followed by the 'then' or 'else' block")
	}
}
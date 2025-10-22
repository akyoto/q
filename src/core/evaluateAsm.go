package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
)

// evaluateAsm converts an assembly instruction to an SSA value.
func (f *Function) evaluateAsm(expr *expression.Expression) (ssa.Value, error) {
	name := expr.Token.StringFrom(f.File.Bytes)

	switch name {
	case "r0":
		return f.Append(&ssa.Register{Register: 0}), nil
	case "r1":
		return f.Append(&ssa.Register{Register: 1}), nil
	case "r2":
		return f.Append(&ssa.Register{Register: 2}), nil
	case "r3":
		return f.Append(&ssa.Register{Register: 3}), nil
	case "r4":
		return f.Append(&ssa.Register{Register: 4}), nil
	case "r5":
		return f.Append(&ssa.Register{Register: 5}), nil
	case "r6":
		return f.Append(&ssa.Register{Register: 6}), nil
	case "r7":
		return f.Append(&ssa.Register{Register: 7}), nil
	case "r8":
		return f.Append(&ssa.Register{Register: 8}), nil
	case "r9":
		return f.Append(&ssa.Register{Register: 9}), nil
	case "r10":
		return f.Append(&ssa.Register{Register: 10}), nil
	case "r11":
		return f.Append(&ssa.Register{Register: 11}), nil
	case "r12":
		return f.Append(&ssa.Register{Register: 12}), nil
	case "r13":
		return f.Append(&ssa.Register{Register: 13}), nil
	case "r14":
		return f.Append(&ssa.Register{Register: 14}), nil
	case "r15":
		return f.Append(&ssa.Register{Register: 15}), nil
	case "sp":
		return f.Append(&ssa.Register{Register: f.CPU.StackPointer}), nil
	}

	return nil, errors.New(&UnknownIdentifier{Name: name}, f.File, expr.Source())
}
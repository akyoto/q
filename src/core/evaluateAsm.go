package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
)

// evaluateAsm converts an assembly instruction to an SSA value.
func (f *Function) evaluateAsm(expr *expression.Expression) (ssa.Value, error) {
	name := expr.Token.String(f.File.Bytes)

	switch name {
	case "r0":
		return f.Append(&ssa.Register{Register: 0}), nil
	case "r1":
		return f.Append(&ssa.Register{Register: 1}), nil
	case "r6":
		return f.Append(&ssa.Register{Register: 6}), nil
	case "r7":
		return f.Append(&ssa.Register{Register: 7}), nil
	case "sp":
		return f.Append(&ssa.Register{Register: f.CPU.StackPointer}), nil
	}

	return nil, errors.New(&UnknownIdentifier{Name: name}, f.File, expr.Source().StartPos)
}
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
	case "sp":
		return f.Append(&ssa.Register{Register: f.CPU.StackPointer}), nil
	}

	return nil, errors.New(&UnknownIdentifier{Name: name}, f.File, expr.Source().StartPos)
}
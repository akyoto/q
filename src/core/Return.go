package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// Return compiles a return instruction.
func (f *Function) Return(tokens token.List) error {
	if len(tokens) == 1 {
		f.Append(&ssa.Return{Source: ssa.Source{StartPos: tokens[0].Position, EndPos: tokens[0].End()}})
		return nil
	}

	expr := expression.Parse(tokens[1:])
	value, err := f.eval(expr)

	if err != nil {
		return err
	}

	f.Append(&ssa.Return{
		Arguments: []ssa.Value{value},
		Source:    ssa.Source(expr.Source()),
	})

	return nil
}
package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// CompileReturn compiles a return instruction.
func (f *Function) CompileReturn(tokens token.List) error {
	expr := expression.Parse(tokens[1:])
	value, err := f.Evaluate(expr)

	if err != nil {
		return err
	}

	f.Append(ssa.Value{
		Type: ssa.Return,
		Args: []*ssa.Value{value},
	})

	return nil
}
package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/ssa"
)

// compileReturn compiles a return instruction.
func (f *Function) compileReturn(node *ast.Return) error {
	if len(node.Values) == 0 {
		f.Append(&ssa.Return{})
		// f.Append(&ssa.Return{Source: ssa.Source{StartPos: tokens[0].Position, EndPos: tokens[0].End()}})
		return nil
	}

	value, err := f.evaluate(node.Values[0])

	if err != nil {
		return err
	}

	f.Append(&ssa.Return{
		Arguments: []ssa.Value{value},
		// Source:    ssa.Source(expr.Source()),
	})

	return nil
}
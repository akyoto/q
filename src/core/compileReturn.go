package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
)

// compileReturn compiles a return instruction.
func (f *Function) compileReturn(node *ast.Return) error {
	if len(node.Values) != len(f.Output) {
		position := node.Token.End()

		if len(node.Values) > 0 {
			position = node.Values[0].Token.Position
		}

		return errors.NewAt(&ReturnCountMismatch{Count: len(node.Values), ExpectedCount: len(f.Output)}, f.File, position)
	}

	if len(node.Values) == 0 {
		f.Append(&ssa.Return{})
		return nil
	}

	args, err := f.decompose(node.Values, f.Output, true)

	if err != nil {
		return err
	}

	f.Append(&ssa.Return{Arguments: args})
	return nil
}
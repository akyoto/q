package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
)

// compileReturn compiles a return instruction.
func (f *Function) compileReturn(node *ast.Return) error {
	if len(node.Values) != len(f.Output) {
		// Special case: Use the multi-return of a single call
		// to satisfy the requirement for multiple return types.
		if len(node.Values) == 1 && len(f.Output) > 1 {
			return f.compileReturnTuple(node)
		}

		position := node.Token.End()

		if len(node.Values) > 0 {
			position = node.Values[0].Source().Start()
		}

		return errors.NewAt(&ReturnCountMismatch{Count: len(node.Values), ExpectedCount: len(f.Output)}, f.File, position)
	}

	if len(node.Values) == 0 {
		f.deleteResources()
		f.Append(&ssa.Return{})
		return nil
	}

	args, err := f.decompose(node.Values, f.Output, true)

	if err != nil {
		return err
	}

	f.deleteResources()
	f.Append(&ssa.Return{Arguments: args})
	return nil
}
package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// compileReturnTuple compiles a return with a single argument
// when multiple output types are expected.
func (f *Function) compileReturnTuple(node *ast.Return) error {
	value, err := f.evaluateRight(node.Values[0])

	if err != nil {
		return err
	}

	tuple, isTuple := value.Type().(*types.Tuple)

	if !isTuple {
		return errors.NewAt(&ReturnCountMismatch{Count: len(node.Values), ExpectedCount: len(f.Output)}, f.File, node.Values[0].Source().Start())
	}

	args, err := f.decomposeTuple(value, tuple, f.Output, node.Values[0].Source())

	if err != nil {
		return err
	}

	if len(args) != len(f.Output) {
		return errors.NewAt(&ReturnCountMismatch{Count: len(args), ExpectedCount: len(f.Output)}, f.File, node.Values[0].Source().Start())
	}

	f.Append(&ssa.Return{Arguments: args})
	return nil
}
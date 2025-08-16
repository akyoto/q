package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// compileReturn compiles a return instruction.
func (f *Function) compileReturn(node *ast.Return) error {
	if len(node.Values) == 0 {
		f.Append(&ssa.Return{})
		return nil
	}

	if len(node.Values) != len(f.Output) {
		return errors.New(&ReturnCountMismatch{Count: len(node.Values), ExpectedCount: len(f.Output)}, f.File, node.Values[0].Token.Position)
	}

	returnValues := make([]ssa.Value, 0, len(node.Values))

	for i, expr := range node.Values {
		value, err := f.evaluate(expr)

		if err != nil {
			return err
		}

		if !types.Is(value.Type(), f.Output[i].Type()) {
			return errors.New(&TypeMismatch{Encountered: value.Type().Name(), Expected: f.Output[i].Type().Name(), ParameterName: f.Output[i].Name, IsReturn: true}, f.File, expr.Token.Position)
		}

		returnValues = append(returnValues, value)
	}

	f.Append(&ssa.Return{Arguments: returnValues})
	return nil
}
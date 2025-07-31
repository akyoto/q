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

	value, err := f.evaluate(node.Values[0])

	if err != nil {
		return err
	}

	if !types.Is(value.Type(), f.Output[0].Type()) {
		return errors.New(&TypeMismatch{Encountered: value.Type().Name(), Expected: f.Output[0].Type().Name(), ParameterName: f.Output[0].Name, IsReturn: true}, f.File, node.Values[0].Token.Position)
	}

	f.Append(&ssa.Return{Arguments: []ssa.Value{value}})
	return nil
}
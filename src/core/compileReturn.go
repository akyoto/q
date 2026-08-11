package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
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
		f.deleteResources(nil)
		f.Append(&ssa.Return{})
		return nil
	}

	values, err := f.evaluateAll(node.Values)

	if err != nil {
		return err
	}

	for i, value := range values {
		given := value.Type()
		expected := f.Output[i].Typ
		_, givenIsResource := given.(*types.Resource)
		expectedResource, expectedIsResource := expected.(*types.Resource)

		if givenIsResource && expectedIsResource {
			f.Block().Unidentify(value)
		}

		if types.Is(given, expected) {
			continue
		}

		if expectedIsResource && types.Is(given, expectedResource.Of) {
			continue
		}

		_, isSyscall := value.(*ssa.Syscall)

		if isSyscall {
			continue
		}

		typeMismatch := &TypeMismatch{
			Encountered:   given.Name(),
			Expected:      expected.Name(),
			ParameterName: f.Output[i].Name,
			IsReturn:      true,
		}

		return errors.New(typeMismatch, f.File, node.Values[i].Source())
	}

	args, err := f.decompose(values)

	if err != nil {
		return err
	}

	f.deleteResources(nil)
	f.Append(&ssa.Return{Arguments: args})
	return nil
}
package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateMethod converts a method expression to an SSA value.
func (f *Function) evaluateMethod(leftValue ssa.Value, left *expression.Expression, right *expression.Expression, expr *expression.Expression) (ssa.Value, error) {
	methodName := right.Token.StringFrom(f.File.Bytes)
	leftUnwrapped := types.Unwrap(leftValue.Type())
	pointer, isPointer := leftUnwrapped.(*types.Pointer)

	if isPointer {
		leftUnwrapped = pointer.To
	}

	call := expr.Parent
	call.Children = append(call.Children, nil)
	copy(call.Children[2:], call.Children[1:])
	call.Children[1] = left

	pkg := f.File.Package
	structure, isStructPointer := leftUnwrapped.(*types.Struct)

	if isStructPointer && structure.Package != "" {
		pkg = structure.Package
	}

	return f.evaluatePackageMember(f.Env.Packages[pkg], methodName, expr)
}
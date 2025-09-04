package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// compileDefinition compiles a define instruction.
func (f *Function) compileDefinition(node *ast.Define) error {
	left := node.Expression.Children[0]
	right := node.Expression.Children[1]
	rightValue, err := f.evaluate(right)

	if err != nil {
		return err
	}

	if left.IsLeaf() {
		call, isCall := rightValue.(*ssa.Call)

		if isCall && len(call.Func.Typ.Output) != 1 {
			return errors.New(&DefinitionCountMismatch{Function: call.Func.String(), Count: 1, ExpectedCount: len(call.Func.Typ.Output)}, f.File, left.Source().StartPos)
		}

		name := left.String(f.File.Bytes)

		if name == "_" {
			return nil
		}

		return f.define(left, rightValue)
	}

	leaves := make([]*expression.Expression, 0, 2)

	for leaf := range left.Leaves() {
		leaves = append(leaves, leaf)
	}

	fn := rightValue.(*ssa.Call).Func

	if len(leaves) != len(fn.Typ.Output) {
		return errors.New(&DefinitionCountMismatch{Function: fn.String(), Count: len(leaves), ExpectedCount: len(fn.Typ.Output)}, f.File, left.Source().StartPos)
	}

	protected := make([]ssa.Value, 0, len(leaves))
	count := 0

	for i, identifier := range leaves {
		name := identifier.String(f.File.Bytes)

		if name == "_" {
			count++
			continue
		}

		_, exists := f.Block().FindIdentifier(name)

		if exists {
			return errors.New(&VariableAlreadyExists{Name: name}, f.File, identifier.Source().StartPos)
		}

		structure, isStructType := types.Unwrap(fn.Typ.Output[i]).(*types.Struct)

		if !isStructType {
			value := &ssa.FromTuple{
				Tuple:  rightValue,
				Index:  count,
				Source: identifier.Source(),
			}

			f.Block().Append(value)
			f.Block().Identify(name, value)

			if value.Type() == types.Error {
				f.Block().Protect(value, protected)
				protected = nil
			} else {
				protected = append(protected, value)
			}

			count++
			continue
		}

		composite := &ssa.Struct{
			Typ:       fn.Typ.Output[i],
			Arguments: make(ssa.Arguments, 0, len(structure.Fields)),
			Source:    identifier.Source(),
		}

		for _, field := range structure.Fields {
			fieldValue := &ssa.FromTuple{
				Tuple:     rightValue,
				Index:     count,
				Structure: composite,
				Source:    identifier.Source(),
			}

			f.Block().Append(fieldValue)
			f.Block().Identify(name+"."+field.Name, fieldValue)
			composite.Arguments = append(composite.Arguments, fieldValue)
			protected = append(protected, fieldValue)
			count++
		}

		f.Block().Identify(name, composite)
		protected = append(protected, composite)
	}

	return nil
}
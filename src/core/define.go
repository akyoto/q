package core

import (
	"fmt"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// define gives a value an identifier.
func (f *Function) define(identifier *expression.Expression, value ssa.Value) error {
	name := identifier.String(f.File.Bytes)
	_, exists := f.Block().FindIdentifier(name)

	if exists {
		return errors.New(&VariableAlreadyExists{Name: name}, f.File, identifier.Source().StartPos)
	}

	// If the value we got was a value that is stored in a variable,
	// it must have been returned from the optimizer as a cached value.
	// We want to assure that every named variable creates a copy of
	// another named variable instead of using the cached value itself
	// because it could lead to incorrect optimizations.
	if f.IsIdentified(value) {
		value = f.copy(value, identifier.Source())
	}

	_, isCall := value.(*ssa.Call)

	if !isCall {
		structure, isStructType := value.(*ssa.Struct)

		if isStructType {
			for i, field := range structure.Typ.Fields {
				f.Block().Identify(fmt.Sprintf("%s.%s", name, field.Name), structure.Arguments[i])
			}
		}

		f.Block().Identify(name, value)
		return nil
	}

	structure, isStructType := value.Type().(*types.Struct)

	if !isStructType {
		f.Block().Identify(name, value)
		return nil
	}

	composite := &ssa.Struct{
		Typ:       structure,
		Arguments: make(ssa.Arguments, 0, len(structure.Fields)),
		Source:    identifier.Source(),
	}

	for i, field := range structure.Fields {
		fieldValue := &ssa.FromTuple{
			Tuple:     value,
			Index:     i,
			Structure: composite,
			Source:    identifier.Source(),
		}

		f.Block().Append(fieldValue)
		f.Block().Identify(fmt.Sprintf("%s.%s", name, field.Name), fieldValue)
		composite.Arguments = append(composite.Arguments, fieldValue)
	}

	f.Block().Identify(name, composite)
	return nil
}
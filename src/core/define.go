package core

import (
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

	for i := range structure.Fields {
		field := &ssa.FromTuple{
			Tuple:     value,
			Index:     i,
			Structure: composite,
			Source:    identifier.Source(),
		}

		f.Block().Append(field)
		composite.Arguments = append(composite.Arguments, field)
	}

	f.Block().Identify(name, composite)
	return nil
}
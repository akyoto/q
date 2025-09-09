package core

import "git.urbach.dev/cli/q/src/ssa"

// copy creates a copy of a value and appends it to the current block.
func (f *Function) copy(value ssa.Value, source ssa.Source) ssa.Value {
	structValue, isStruct := value.(*ssa.Struct)

	if isStruct {
		c := &ssa.Struct{
			Typ:       structValue.Typ,
			Arguments: make(ssa.Arguments, len(structValue.Arguments)),
		}

		for i, field := range structValue.Arguments {
			c.Arguments[i] = f.copy(field, source)
		}

		return c
	}

	c := &ssa.Copy{
		Value:  value,
		Typ:    value.Type(),
		Source: source,
	}

	f.Block().Append(c)
	return c
}
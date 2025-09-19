package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
)

// evaluateRight creates a load from memory if the value
// was a memory address, otherwise returns the raw value.
func (f *Function) evaluateRight(expr *expression.Expression) (ssa.Value, error) {
	value, err := f.evaluate(expr)

	if err != nil {
		return nil, err
	}

	memory, isMemory := value.(*ssa.Memory)

	if isMemory {
		load := f.Append(&ssa.Load{
			Memory: memory,
			Source: memory.Source,
		})

		return load, nil
	}

	return value, nil
}
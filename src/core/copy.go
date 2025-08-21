package core

import "git.urbach.dev/cli/q/src/ssa"

// copy creates a copy of a value and appends it to the current block.
func (f *Function) copy(value ssa.Value, source ssa.Source) *ssa.Copy {
	c := &ssa.Copy{
		Value:  value,
		Source: source,
	}

	f.Block().Append(c)
	return c
}
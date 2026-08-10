package core

import "git.urbach.dev/cli/q/src/ssa"

// findIdentifier returns the value for an identifier,
// whether it exists and whether it's partially defined.
func (f *Function) findIdentifier(name string) (ssa.Value, bool, bool) {
	value, exists := f.Block().FindIdentifier(name)
	phi, isPhi := value.(*ssa.Phi)
	partial := isPhi && phi.IsPartiallyUndefined()
	return value, exists, partial
}
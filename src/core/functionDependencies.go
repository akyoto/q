package core

import "git.urbach.dev/cli/q/src/set"

// functionDependencies contains references to other functions.
type functionDependencies struct {
	Calls   set.Ordered[*Function]
	Globals set.Ordered[*Global]
}

// IsLeaf returns true if the function doesn't call other functions.
func (f *functionDependencies) IsLeaf() bool {
	return f.Calls.Count() == 0
}
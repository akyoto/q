package core

import "git.urbach.dev/cli/q/src/set"

// functionDependencies contains references to other functions.
type functionDependencies struct {
	Dependencies set.Ordered[*Function]
}

// IsLeaf returns true if the function doesn't call other functions.
func (f *functionDependencies) IsLeaf() bool {
	return f.Dependencies.Count() == 0
}
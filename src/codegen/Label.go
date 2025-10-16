package codegen

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// Label represents a jump target.
type Label struct {
	ssa.Void
	Name string
}

// Equals returns true if both labels have the same name.
func (a *Label) Equals(v ssa.Value) bool {
	b, sameType := v.(*Label)

	if !sameType {
		return false
	}

	return a.Name == b.Name
}

// Inputs returns nil because a label does not have dependencies.
func (v *Label) Inputs() []ssa.Value {
	return nil
}

// IsPure returns false because a label is not a constant.
func (v *Label) IsPure() bool {
	return false
}

// Replace does nothing on a label.
func (v *Label) Replace(ssa.Value, ssa.Value) {}

// String returns the human-readable "clean" version
// of the label without the function prefix.
func (v *Label) String() string {
	return ssa.CleanLabel(v.Name)
}

// Type returns the Void type for labels.
func (v *Label) Type() types.Type {
	return types.Void
}
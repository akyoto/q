package ssa

import "git.urbach.dev/cli/q/src/types"

// Void is the absence of a value.
type Void struct{}

// AddUser does nothing because a void value cannot be used as a dependency.
func (v *Void) AddUser(Value) {}

// IsConst returns false because a void value is not a constant.
func (v *Void) IsConst() bool { return false }

// RemoveUser does nothing because a void value cannot be used as a dependency.
func (v *Void) RemoveUser(Value) {}

// Type returns the void type.
func (v *Void) Type() types.Type { return types.Void }

// Users returns nil because a void value has no users.
func (v *Void) Users() []Value { return nil }